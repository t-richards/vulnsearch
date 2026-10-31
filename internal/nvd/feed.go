package nvd

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"os"
	"strings"
	"time"
)

// StatusRejected marks CVEs which have been withdrawn.
const StatusRejected = "Rejected"

// Errors for malformed feeds.
var (
	errNoVulnerabilities = errors.New("feed has no vulnerabilities")
	errUnexpectedToken   = errors.New("unexpected token")
)

// Time is a timestamp in the feed's format, which omits the time zone (UTC).
type Time struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Time) UnmarshalJSON(buf []byte) error {
	var s string

	err := json.Unmarshal(buf, &s)
	if err != nil {
		return err //nolint:wrapcheck // Only called by encoding/json.
	}

	t.Time, err = time.Parse("2006-01-02T15:04:05.999999999", s)
	if err != nil {
		return fmt.Errorf("parsing time: %w", err)
	}

	return nil
}

// CVE is a single entry in the feed's vulnerabilities array. Only the fields
// used by vulnsearch are decoded.
type CVE struct {
	ID             string          `json:"id"`
	Published      Time            `json:"published"`
	LastModified   Time            `json:"lastModified"`
	VulnStatus     string          `json:"vulnStatus"`
	Descriptions   []LangString    `json:"descriptions"`
	Metrics        Metrics         `json:"metrics"`
	Weaknesses     []Weakness      `json:"weaknesses"`
	Configurations []Configuration `json:"configurations"`
}

// LangString is a string in a given language.
type LangString struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

// Weakness lists the CWE IDs a source assigned to a CVE.
type Weakness struct {
	Source      string       `json:"source"`
	Type        string       `json:"type"`
	Description []LangString `json:"description"`
}

// Configuration is a set of nodes, combined with Operator (AND/OR), describing
// the platforms a CVE applies to.
type Configuration struct {
	Operator string `json:"operator"`
	Negate   bool   `json:"negate"`
	Nodes    []Node `json:"nodes"`
}

// Node is a set of CPE match criteria combined with Operator.
type Node struct {
	Operator string     `json:"operator"`
	Negate   bool       `json:"negate"`
	CPEMatch []CPEMatch `json:"cpeMatch"`
}

// CPEMatch is a CPE name with optional version range bounds. Non-vulnerable
// matches describe the platform a vulnerable product must run on.
type CPEMatch struct {
	Vulnerable            bool   `json:"vulnerable"`
	Criteria              string `json:"criteria"`
	VersionStartIncluding string `json:"versionStartIncluding"`
	VersionStartExcluding string `json:"versionStartExcluding"`
	VersionEndIncluding   string `json:"versionEndIncluding"`
	VersionEndExcluding   string `json:"versionEndExcluding"`
}

// Metrics holds the CVSS scores from each source, per CVSS version.
type Metrics struct {
	V40 []Metric `json:"cvssMetricV40"`
	V31 []Metric `json:"cvssMetricV31"`
	V30 []Metric `json:"cvssMetricV30"`
	V2  []Metric `json:"cvssMetricV2"`
}

// Metric is a CVSS score from a single source. Type is "Primary" for NVD's
// own analysis, and "Secondary" for other sources such as the CNA.
type Metric struct {
	Source   string   `json:"source"`
	Type     string   `json:"type"`
	CVSSData CVSSData `json:"cvssData"`
}

// CVSSData holds the CVSS vector and base score.
type CVSSData struct {
	Version      string  `json:"version"`
	VectorString string  `json:"vectorString"`
	BaseScore    float64 `json:"baseScore"`
}

// Description returns the English description of the CVE.
func (c CVE) Description() string {
	for _, d := range c.Descriptions {
		if d.Lang == "en" {
			return d.Value
		}
	}

	return ""
}

// CWEIDs returns the distinct CWE IDs assigned to the CVE, comma-separated.
func (c CVE) CWEIDs() string {
	var ids []string

	seen := map[string]bool{}

	for _, w := range c.Weaknesses {
		for _, d := range w.Description {
			if d.Lang == "en" && !seen[d.Value] {
				seen[d.Value] = true
				ids = append(ids, d.Value)
			}
		}
	}

	return strings.Join(ids, ", ")
}

// V4Score returns the CVSS v4.0 base score, if any.
func (m Metrics) V4Score() (float64, bool) {
	return preferredScore(m.V40)
}

// V3Score returns the CVSS v3.1 base score, falling back to v3.0, if any.
func (m Metrics) V3Score() (float64, bool) {
	score, ok := preferredScore(m.V31)
	if ok {
		return score, true
	}

	return preferredScore(m.V30)
}

// V2Score returns the CVSS v2.0 base score, if any.
func (m Metrics) V2Score() (float64, bool) {
	return preferredScore(m.V2)
}

// preferredScore picks NVD's primary score, else the first secondary score.
func preferredScore(metrics []Metric) (float64, bool) {
	for _, m := range metrics {
		if m.Type == "Primary" {
			return m.CVSSData.BaseScore, true
		}
	}

	if len(metrics) > 0 {
		return metrics[0].CVSSData.BaseScore, true
	}

	return 0, false
}

// VulnerableMatches returns the distinct CPE matches which identify vulnerable
// products. Platform-only (non-vulnerable) and negated matches are skipped.
func (c CVE) VulnerableMatches() []CPEMatch {
	var matches []CPEMatch

	seen := map[CPEMatch]bool{}

	for _, config := range c.Configurations {
		if config.Negate {
			continue
		}

		for _, node := range config.Nodes {
			if node.Negate {
				continue
			}

			for _, match := range node.CPEMatch {
				if match.Vulnerable && !seen[match] {
					seen[match] = true
					matches = append(matches, match)
				}
			}
		}
	}

	return matches
}

// OpenArchive decodes the gzipped feed at path. See Decode.
func OpenArchive(path string) iter.Seq2[CVE, error] {
	return func(yield func(CVE, error) bool) {
		file, err := os.Open(path)
		if err != nil {
			yield(CVE{}, fmt.Errorf("opening archive: %w", err))

			return
		}
		defer file.Close()

		reader, err := gzip.NewReader(file)
		if err != nil {
			yield(CVE{}, fmt.Errorf("decompressing %s: %w", path, err))

			return
		}
		defer reader.Close()

		for cve, err := range Decode(reader) {
			if !yield(cve, err) || err != nil {
				return
			}
		}
	}
}

// Decode streams the CVEs from a JSON feed one at a time, so the (potentially
// several hundred megabyte) feed is never held in memory all at once.
// Iteration stops after the first error.
func Decode(r io.Reader) iter.Seq2[CVE, error] {
	return func(yield func(CVE, error) bool) {
		dec := json.NewDecoder(r)

		err := seekVulnerabilities(dec)
		if err != nil {
			yield(CVE{}, err)

			return
		}

		for dec.More() {
			var item struct {
				CVE CVE `json:"cve"`
			}

			err := dec.Decode(&item)
			if err != nil {
				yield(CVE{}, fmt.Errorf("decoding CVE: %w", err))

				return
			}

			if !yield(item.CVE, nil) {
				return
			}
		}
	}
}

// seekVulnerabilities advances dec to the first element of the top-level
// "vulnerabilities" array, skipping over any other keys.
func seekVulnerabilities(dec *json.Decoder) error {
	tok, err := dec.Token()
	if err != nil {
		return fmt.Errorf("reading feed start: %w", err)
	}

	if tok != json.Delim('{') {
		return fmt.Errorf("reading feed start: %w %v, expected object", errUnexpectedToken, tok)
	}

	for dec.More() {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf("reading feed key: %w", err)
		}

		if tok == "vulnerabilities" {
			tok, err := dec.Token()
			if err != nil {
				return fmt.Errorf("reading vulnerabilities start: %w", err)
			}

			if tok != json.Delim('[') {
				return fmt.Errorf("reading vulnerabilities start: %w %v, expected array", errUnexpectedToken, tok)
			}

			return nil
		}

		// Skip the value of any other key.
		var skip json.RawMessage

		err = dec.Decode(&skip)
		if err != nil {
			return fmt.Errorf("skipping feed key %v: %w", tok, err)
		}
	}

	return errNoVulnerabilities
}
