package cpe

import (
	"cmp"
	"strings"
	"unicode"
)

// Match is the version constraint of an NVD CPE match criterion. Version is
// either an exact version, Any (optionally bounded by the range fields), or
// NotApplicable. Empty range fields are unbounded.
type Match struct {
	Version        string
	StartIncluding string
	StartExcluding string
	EndIncluding   string
	EndExcluding   string
}

// Contains reports whether version satisfies the constraint.
func (m Match) Contains(version string) bool {
	switch m.Version {
	case Any, "":
		// Wildcard; fall through to the range checks below.
	case NotApplicable:
		return version == NotApplicable
	default:
		return Compare(version, m.Version) == 0
	}

	return m.inRange(version)
}

// String describes the constraint, e.g. "2.4.0 ≤ v < 2.4.50" or "< 2.4.66".
// It never contains a comma, so lists of constraints can be comma-separated.
func (m Match) String() string {
	switch m.Version {
	case Any, "":
	case NotApplicable:
		return "n/a"
	default:
		return m.Version
	}

	start, startOp := m.StartIncluding, "≤"
	if start == "" {
		start, startOp = m.StartExcluding, "<"
	}

	end, endOp := m.EndIncluding, "≤"
	if end == "" {
		end, endOp = m.EndExcluding, "<"
	}

	switch {
	case start != "" && end != "":
		return start + " " + startOp + " v " + endOp + " " + end
	case start != "":
		return strings.NewReplacer("≤", "≥", "<", ">").Replace(startOp) + " " + start
	case end != "":
		return endOp + " " + end
	}

	return "all versions"
}

// CompareMatches orders matches by the lowest version they include, so that
// exact versions and ranges sort together.
func CompareMatches(a, b Match) int {
	c := Compare(a.lowerBound(), b.lowerBound())
	if c != 0 {
		return c
	}

	return strings.Compare(a.String(), b.String())
}

// lowerBound is the lowest version the constraint includes, or "" if unbounded.
func (m Match) lowerBound() string {
	switch m.Version {
	case Any, "":
		return cmp.Or(m.StartIncluding, m.StartExcluding)
	default:
		return m.Version
	}
}

// inRange reports whether version is within the (possibly unbounded) range.
func (m Match) inRange(version string) bool {
	switch {
	case m.StartIncluding != "" && Compare(version, m.StartIncluding) < 0:
		return false
	case m.StartExcluding != "" && Compare(version, m.StartExcluding) <= 0:
		return false
	case m.EndIncluding != "" && Compare(version, m.EndIncluding) > 0:
		return false
	case m.EndExcluding != "" && Compare(version, m.EndExcluding) >= 0:
		return false
	}

	return true
}

// preReleaseRank orders well-known pre-release labels. These sort before the
// release they precede ("1.0rc1" < "1.0"); any other letters are treated as
// a post-release suffix ("1.0.2" < "1.0.2k"), as vendors like OpenSSL use them.
var preReleaseRank = map[string]int{
	"dev":      1,
	"snapshot": 1,
	"alpha":    2,
	"beta":     3,
	"pre":      4,
	"preview":  4,
	"rc":       5,
}

// Compare compares two vendor version strings, returning -1, 0 or +1.
//
// Versions in the wild follow no single scheme, so this is a best-effort
// ordering: strings are split into runs of digits and runs of letters,
// separators are ignored, digit runs compare numerically, letter runs
// compare alphabetically (pre-release labels sort lowest), and missing
// trailing components count as zero.
func Compare(a, b string) int {
	at, bt := tokenize(a), tokenize(b)

	for i := range max(len(at), len(bt)) {
		switch {
		case i >= len(at) && isZero(bt[i]), i >= len(bt) && isZero(at[i]):
			continue // "1.0" == "1.0.0"
		case i >= len(at):
			return -tail(bt[i])
		case i >= len(bt):
			return tail(at[i])
		}

		c := compareTokens(at[i], bt[i])
		if c != 0 {
			return c
		}
	}

	return 0
}

// tail gives the sign for a version which has the extra token t compared to an
// otherwise-equal version: pre-release labels make it older, anything else newer.
func tail(t token) int {
	if !t.numeric && preReleaseRank[t.text] > 0 {
		return -1
	}

	return 1
}

func isZero(t token) bool {
	return t.numeric && strings.Trim(t.text, "0") == ""
}

type token struct {
	text    string
	numeric bool
}

func compareTokens(a, b token) int {
	switch {
	case a.numeric && b.numeric:
		return compareNumeric(a.text, b.text)
	case a.numeric:
		return 1
	case b.numeric:
		return -1
	}

	ar, br := preReleaseRank[a.text], preReleaseRank[b.text]

	switch {
	case ar > 0 && br > 0:
		return sign(ar - br)
	case ar > 0:
		return -1
	case br > 0:
		return 1
	}

	return strings.Compare(a.text, b.text)
}

// compareNumeric compares digit strings of any length without overflow.
func compareNumeric(a, b string) int {
	a = strings.TrimLeft(a, "0")
	b = strings.TrimLeft(b, "0")

	if len(a) != len(b) {
		return sign(len(a) - len(b))
	}

	return strings.Compare(a, b)
}

func tokenize(s string) []token {
	var tokens []token

	s = strings.ToLower(s)
	start := -1

	flush := func(end int) {
		if start >= 0 {
			tokens = append(tokens, token{text: s[start:end], numeric: isDigit(rune(s[start]))})
			start = -1
		}
	}

	for i, r := range s {
		switch {
		case !unicode.IsLetter(r) && !isDigit(r):
			flush(i)
		case start >= 0 && isDigit(r) != isDigit(rune(s[start])):
			flush(i)

			start = i
		case start < 0:
			start = i
		}
	}

	flush(len(s))

	return tokens
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func sign(n int) int {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return 1
	}

	return 0
}
