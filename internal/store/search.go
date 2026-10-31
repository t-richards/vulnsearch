package store

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/t-richards/vulnsearch/internal/cpe"
	"github.com/t-richards/vulnsearch/internal/db"
)

// suggestionLimit caps the number of search suggestions returned.
const suggestionLimit = 100

// likeEscaper escapes LIKE wildcards; "_" is common in CPE product names.
var likeEscaper = strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)

func likePrefix(s string) string {
	return likeEscaper.Replace(s) + "%"
}

// Vendors suggests vendors starting with prefix.
func (s *Store) Vendors(ctx context.Context, prefix string) ([]string, error) {
	vendors, err := s.SearchVendors(ctx, likePrefix(prefix))
	if err != nil {
		return nil, fmt.Errorf("while searching vendors: %w", err)
	}

	return vendors, nil
}

// Products suggests a vendor's products starting with prefix.
func (s *Store) Products(ctx context.Context, vendor, prefix string) ([]string, error) {
	products, err := s.SearchProducts(ctx, db.SearchProductsParams{
		Vendor:  vendor,
		Product: likePrefix(prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("while searching products: %w", err)
	}

	return products, nil
}

// Versions suggests a product's versions starting with prefix, in version order.
func (s *Store) Versions(ctx context.Context, vendor, product, prefix string) ([]string, error) {
	versions, err := s.SearchVersions(ctx, db.SearchVersionsParams{
		Vendor:  vendor,
		Product: product,
		Version: likePrefix(prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("while searching versions: %w", err)
	}

	// Sorted here, as SQL can only sort these lexically ("10.0" < "9.0").
	slices.SortFunc(versions, cpe.Compare)

	return versions[:min(len(versions), suggestionLimit)], nil
}

// AffectedCVE is a CVE which affects a product, along with the match criteria
// which apply to the product version.
type AffectedCVE struct {
	db.CVE

	Matches []cpe.Match // Distinct and sorted.
}

// AffectedVersions describes the matching versions, e.g. "2.4.4, 2.4.5".
func (c AffectedCVE) AffectedVersions() string {
	versions := make([]string, len(c.Matches))
	for i, match := range c.Matches {
		versions[i] = match.String()
	}

	return strings.Join(versions, ", ")
}

// AffectedCVEs finds the CVEs affecting a version of a product, most recently
// published first. An empty version finds CVEs affecting any version.
func (s *Store) AffectedCVEs(ctx context.Context, vendor, product, version string) ([]AffectedCVE, error) {
	rows, err := s.ListMatchesForProduct(ctx, db.ListMatchesForProductParams{
		Vendor:  vendor,
		Product: product,
	})
	if err != nil {
		return nil, fmt.Errorf("while listing matches: %w", err)
	}

	var cves []AffectedCVE

	index := map[string]int{}

	for _, row := range rows {
		match := cpe.Match{
			Version:        row.CpeMatch.Version,
			StartIncluding: row.CpeMatch.VersionStartIncluding,
			StartExcluding: row.CpeMatch.VersionStartExcluding,
			EndIncluding:   row.CpeMatch.VersionEndIncluding,
			EndExcluding:   row.CpeMatch.VersionEndExcluding,
		}

		if version != "" && !match.Contains(version) {
			continue
		}

		i, ok := index[row.CVE.ID]
		if !ok {
			i = len(cves)
			index[row.CVE.ID] = i
			cves = append(cves, AffectedCVE{CVE: row.CVE})
		}

		cves[i].Matches = append(cves[i].Matches, match)
	}

	// Criteria can repeat once the CPE fields we don't store are dropped, e.g.
	// one per patch level ("2.4.4:p1", "2.4.4:p2", ...) all become "2.4.4".
	for i := range cves {
		slices.SortFunc(cves[i].Matches, cpe.CompareMatches)
		cves[i].Matches = slices.Compact(cves[i].Matches)
	}

	return cves, nil
}
