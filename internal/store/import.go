package store

import (
	"context"
	"database/sql"
	"fmt"
	"iter"
	"log"

	"github.com/t-richards/vulnsearch/internal/cpe"
	"github.com/t-richards/vulnsearch/internal/db"
	"github.com/t-richards/vulnsearch/internal/nvd"
)

// Import loads CVEs into the database in a single transaction, replacing any
// existing data for each CVE and removing rejected CVEs. It returns the number
// of CVEs processed. Call RebuildProducts once all imports are done.
func (s *Store) Import(ctx context.Context, cves iter.Seq2[nvd.CVE, error]) (int, error) {
	tx, qtx, err := s.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback() //nolint:errcheck

	count := 0

	for cve, err := range cves {
		if err != nil {
			return count, err
		}

		err = importCVE(ctx, qtx, cve)
		if err != nil {
			return count, fmt.Errorf("while importing %s: %w", cve.ID, err)
		}

		count++
	}

	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("while committing import: %w", err)
	}

	return count, nil
}

func importCVE(ctx context.Context, q *db.Queries, cve nvd.CVE) error {
	if cve.VulnStatus == nvd.StatusRejected {
		return q.DeleteCVE(ctx, cve.ID)
	}

	v4, v4ok := cve.Metrics.V4Score()
	v3, v3ok := cve.Metrics.V3Score()
	v2, v2ok := cve.Metrics.V2Score()

	err := q.UpsertCVE(ctx, db.UpsertCVEParams{
		ID:           cve.ID,
		Description:  cve.Description(),
		CWEIDs:       cve.CWEIDs(),
		CVSSv4Score:  sql.NullFloat64{Float64: v4, Valid: v4ok},
		CVSSv3Score:  sql.NullFloat64{Float64: v3, Valid: v3ok},
		CVSSv2Score:  sql.NullFloat64{Float64: v2, Valid: v2ok},
		Published:    cve.Published.Time,
		LastModified: cve.LastModified.Time,
	})
	if err != nil {
		return fmt.Errorf("while saving CVE: %w", err)
	}

	err = q.DeleteMatchesForCVE(ctx, cve.ID)
	if err != nil {
		return fmt.Errorf("while clearing matches: %w", err)
	}

	for _, match := range cve.VulnerableMatches() {
		name, err := cpe.Parse(match.Criteria)
		if err != nil {
			// One malformed criterion shouldn't sink the whole import.
			log.Printf("Skipping match for %s: %v", cve.ID, err)

			continue
		}

		err = q.CreateMatch(ctx, db.CreateMatchParams{
			CVEID:                 cve.ID,
			Vendor:                name.Vendor,
			Product:               name.Product,
			Version:               name.Version,
			VersionStartIncluding: match.VersionStartIncluding,
			VersionStartExcluding: match.VersionStartExcluding,
			VersionEndIncluding:   match.VersionEndIncluding,
			VersionEndExcluding:   match.VersionEndExcluding,
		})
		if err != nil {
			return fmt.Errorf("while saving match: %w", err)
		}
	}

	return nil
}

// RebuildProducts regenerates the vendor, product and version search
// suggestions from the imported CVE data.
func (s *Store) RebuildProducts(ctx context.Context) error {
	tx, qtx, err := s.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback() //nolint:errcheck

	for _, step := range []func(context.Context) error{
		qtx.ClearProducts,
		qtx.ClearProductVersions,
		qtx.RebuildProducts,
		qtx.RebuildProductVersions,
	} {
		err = step(ctx)
		if err != nil {
			return fmt.Errorf("while rebuilding products: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("while committing products: %w", err)
	}

	return nil
}
