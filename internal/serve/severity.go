package serve

import (
	"database/sql"

	"github.com/t-richards/vulnsearch/internal/store"
)

// severityBand is a qualitative severity rating for scores >= min.
type severityBand struct {
	name  string
	scale string
	min   float64
}

// CVSS v3.x and v4.0 share qualitative severity ratings.
var cvssBands = []severityBand{
	{"Critical", "9.0–10.0", 9.0},
	{"High", "7.0–8.9", 7.0},
	{"Medium", "4.0–6.9", 4.0},
	{"Low", "0.1–3.9", 0.1},
	{"None", "0.0", 0},
}

var cvssV2Bands = []severityBand{
	{"High", "7.0–10.0", 7.0},
	{"Medium", "4.0–6.9", 4.0},
	{"Low", "0.0–3.9", 0},
}

// severityTable counts CVEs by severity for one CVSS version.
type severityTable struct {
	Title string
	Rows  []severityRow
	Total int
}

type severityRow struct {
	Severity string
	Scale    string
	Count    int
}

func severityTables(cves []store.AffectedCVE) []severityTable {
	v4 := make([]sql.NullFloat64, len(cves))
	v3 := make([]sql.NullFloat64, len(cves))
	v2 := make([]sql.NullFloat64, len(cves))

	for i, cve := range cves {
		v4[i], v3[i], v2[i] = cve.CVSSv4Score, cve.CVSSv3Score, cve.CVSSv2Score
	}

	return []severityTable{
		newSeverityTable("CVSS v4.0", cvssBands, v4),
		newSeverityTable("CVSS v3.x", cvssBands, v3),
		newSeverityTable("CVSS v2.0", cvssV2Bands, v2),
	}
}

// newSeverityTable counts scores into bands. CVEs lacking a score for this CVSS
// version are counted separately, rather than as a score of zero.
func newSeverityTable(title string, bands []severityBand, scores []sql.NullFloat64) severityTable {
	rows := make([]severityRow, len(bands), len(bands)+1)
	for i, band := range bands {
		rows[i] = severityRow{Severity: band.name, Scale: band.scale}
	}

	rows = append(rows, severityRow{Severity: "Unscored", Scale: "—"})
	unscored := &rows[len(rows)-1]

	for _, score := range scores {
		if !score.Valid {
			unscored.Count++

			continue
		}

		for i, band := range bands {
			if score.Float64 >= band.min {
				rows[i].Count++

				break
			}
		}
	}

	return severityTable{Title: title, Rows: rows, Total: len(scores)}
}
