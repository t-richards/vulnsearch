package nvd_test

import (
	"os"
	"strings"
	"testing"

	"github.com/matryer/is"

	"github.com/t-richards/vulnsearch/internal/nvd"
)

func decodeFixture(t *testing.T) []nvd.CVE {
	t.Helper()

	is := is.New(t)

	file, err := os.Open("testdata/feed.json")
	is.NoErr(err)

	defer file.Close()

	var cves []nvd.CVE

	for cve, err := range nvd.Decode(file) {
		is.NoErr(err)

		cves = append(cves, cve)
	}

	return cves
}

func TestDecode(t *testing.T) {
	is := is.New(t)

	cves := decodeFixture(t)
	is.Equal(len(cves), 3)

	// A CVE with a version range.
	iterm := cves[0]
	is.Equal(iterm.ID, "CVE-2025-22275")
	is.Equal(iterm.Published.Format("2006-01-02T15:04:05.000"), "2025-01-03T05:15:08.243")
	is.Equal(iterm.CWEIDs(), "CWE-532")
	is.True(strings.HasPrefix(iterm.Description(), "iTerm2"))

	score, ok := iterm.Metrics.V3Score()
	is.True(ok)
	is.Equal(score, 9.3)

	_, ok = iterm.Metrics.V2Score()
	is.True(!ok) // No v2 score is distinct from a score of zero.

	is.Equal(iterm.VulnerableMatches(), []nvd.CPEMatch{{
		Vulnerable:            true,
		Criteria:              "cpe:2.3:a:iterm2:iterm2:*:*:*:*:*:*:*:*",
		VersionStartIncluding: "3.5.6",
		VersionEndExcluding:   "3.5.11",
	}})

	// A CVE which only affects a product running on a (non-vulnerable) platform.
	bible := cves[1]
	is.Equal(bible.ID, "CVE-2010-0157")

	matches := bible.VulnerableMatches()
	is.Equal(len(matches), 1)
	is.Equal(matches[0].Criteria, "cpe:2.3:a:joomlabiblestudy:com_biblestudy:6.1:*:*:*:*:*:*:*")

	// A rejected CVE.
	is.Equal(cves[2].VulnStatus, nvd.StatusRejected)
}

func TestDecodeInvalid(t *testing.T) {
	for _, in := range []string{``, `[]`, `{"format": "NVD_CVE"}`, `{"vulnerabilities": {}}`, `{"vulnerabilities": [{]}`} {
		t.Run(in, func(t *testing.T) {
			is := is.New(t)

			var err error
			for _, err = range nvd.Decode(strings.NewReader(in)) {
				if err != nil {
					break
				}
			}

			is.True(err != nil)
		})
	}
}
