package cpe_test

import (
	"testing"

	"github.com/matryer/is"

	"github.com/t-richards/vulnsearch/internal/cpe"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"2.4.49", "2.4.49", 0},
		{"2.4.49", "2.4.5", 1},
		{"2.4.9", "2.4.10", -1},
		{"1.0", "1.0.0", 0},
		{"1.0", "1.0.1", -1},
		{"1.0.2", "1.0.2k", -1},
		{"1.0.2k", "1.0.2zf", -1},
		{"1.0rc1", "1.0", -1},
		{"1.0-beta2", "1.0-rc1", -1},
		{"1.0.0-alpha", "1.0.0-beta", -1},
		{"5.10", "5.10.0-rc3", 1},
		{"8.0_update291", "8.0 update291", 0},
		{"10.0.19041.1", "10.0.19041.1023", -1},
		{"007", "7", 0},
		{"99999999999999999999", "100000000000000000000", -1},
		{"V2.0", "v2.0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.a+" vs "+tt.b, func(t *testing.T) {
			is := is.New(t)
			is.Equal(cpe.Compare(tt.a, tt.b), tt.want)
			is.Equal(cpe.Compare(tt.b, tt.a), -tt.want) // antisymmetric
		})
	}
}

func TestMatchContains(t *testing.T) {
	is := is.New(t)

	// Exact version.
	exact := cpe.Match{Version: "2.4.49"}
	is.True(exact.Contains("2.4.49"))
	is.True(!exact.Contains("2.4.50"))

	// Wildcard with no bounds matches everything.
	all := cpe.Match{Version: cpe.Any}
	is.True(all.Contains("0.1"))
	is.True(all.Contains("99"))

	// Half-open range [2.4.0, 2.4.50).
	r := cpe.Match{Version: cpe.Any, StartIncluding: "2.4.0", EndExcluding: "2.4.50"}
	is.True(r.Contains("2.4.0"))
	is.True(r.Contains("2.4.49"))
	is.True(!r.Contains("2.4.50"))
	is.True(!r.Contains("2.2.34"))

	// Closed / open bounds.
	r = cpe.Match{Version: cpe.Any, StartExcluding: "1.0", EndIncluding: "1.5"}
	is.True(!r.Contains("1.0"))
	is.True(r.Contains("1.0.1"))
	is.True(r.Contains("1.5"))
	is.True(!r.Contains("1.5.1"))

	// Not applicable only matches itself.
	na := cpe.Match{Version: cpe.NotApplicable}
	is.True(na.Contains("-"))
	is.True(!na.Contains("1.0"))
}

func TestMatchString(t *testing.T) {
	is := is.New(t)

	is.Equal(cpe.Match{Version: "2.4.49"}.String(), "2.4.49")
	is.Equal(cpe.Match{Version: cpe.Any}.String(), "all versions")
	is.Equal(cpe.Match{Version: cpe.NotApplicable}.String(), "n/a")
	is.Equal(cpe.Match{Version: cpe.Any, StartIncluding: "2.4.0", EndExcluding: "2.4.50"}.String(), ">= 2.4.0, < 2.4.50")
}
