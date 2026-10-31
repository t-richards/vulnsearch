package store_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"

	"github.com/t-richards/vulnsearch/internal/nvd"
	"github.com/t-richards/vulnsearch/internal/store"
)

// openFixture opens a fresh database loaded with the nvd package's test feed.
func openFixture(t *testing.T) *store.Store {
	t.Helper()

	is := is.New(t)

	st, err := store.Open(t.Context(), filepath.Join(t.TempDir(), "test.sqlite3"))
	is.NoErr(err)
	t.Cleanup(func() { st.Close() })

	file, err := os.Open("../nvd/testdata/feed.json")
	is.NoErr(err)

	defer file.Close()

	count, err := st.Import(t.Context(), nvd.Decode(file))
	is.NoErr(err)
	is.Equal(count, 3)

	is.NoErr(st.RebuildProducts(t.Context()))

	return st
}

func TestAffectedCVEsRange(t *testing.T) {
	is := is.New(t)
	st := openFixture(t)
	ctx := t.Context()

	// iTerm2 >= 3.5.6, < 3.5.11 is affected.
	for version, want := range map[string]int{
		"3.5.5":  0,
		"3.5.6":  1,
		"3.5.10": 1,
		"3.5.11": 0,
		"":       1, // Any version.
	} {
		cves, err := st.AffectedCVEs(ctx, "iterm2", "iterm2", version)
		is.NoErr(err)
		is.Equal(len(cves), want) // version
	}

	cves, err := st.AffectedCVEs(ctx, "ITerm2", "iTerm2", "3.5.10") // Case insensitive.
	is.NoErr(err)
	is.Equal(len(cves), 1)
	is.Equal(cves[0].ID, "CVE-2025-22275")
	is.Equal(cves[0].CVSSv3Score.Float64, 9.3)
	is.True(!cves[0].CVSSv2Score.Valid)
	is.Equal(cves[0].Matches[0].String(), ">= 3.5.6, < 3.5.11")
}

func TestAffectedCVEsSkipsPlatforms(t *testing.T) {
	is := is.New(t)
	st := openFixture(t)
	ctx := t.Context()

	// The vulnerable component is linked...
	cves, err := st.AffectedCVEs(ctx, "joomlabiblestudy", "com_biblestudy", "6.1")
	is.NoErr(err)
	is.Equal(len(cves), 1)

	// ...but the platform it runs on is not.
	cves, err = st.AffectedCVEs(ctx, "joomla", "joomla!", "")
	is.NoErr(err)
	is.Equal(len(cves), 0)
}

func TestReimportAndReject(t *testing.T) {
	is := is.New(t)
	st := openFixture(t)
	ctx := t.Context()

	// Re-importing replaces matches rather than duplicating them.
	file, err := os.Open("../nvd/testdata/feed.json")
	is.NoErr(err)

	defer file.Close()

	_, err = st.Import(ctx, nvd.Decode(file))
	is.NoErr(err)

	cves, err := st.AffectedCVEs(ctx, "iterm2", "iterm2", "3.5.10")
	is.NoErr(err)
	is.Equal(len(cves), 1)
	is.Equal(len(cves[0].Matches), 1)

	// Rejecting a CVE removes it, and its matches.
	reject := func(yield func(nvd.CVE, error) bool) {
		yield(nvd.CVE{ID: "CVE-2025-22275", VulnStatus: nvd.StatusRejected}, nil)
	}

	_, err = st.Import(ctx, reject)
	is.NoErr(err)

	cves, err = st.AffectedCVEs(ctx, "iterm2", "iterm2", "")
	is.NoErr(err)
	is.Equal(len(cves), 0)
}

func TestSuggestions(t *testing.T) {
	is := is.New(t)
	st := openFixture(t)
	ctx := t.Context()

	vendors, err := st.Vendors(ctx, "IT")
	is.NoErr(err)
	is.Equal(vendors, []string{"iterm2"})

	// LIKE wildcards in the prefix are matched literally.
	products, err := st.Products(ctx, "joomlabiblestudy", "com_")
	is.NoErr(err)
	is.Equal(products, []string{"com_biblestudy"})

	products, err = st.Products(ctx, "joomlabiblestudy", "co%")
	is.NoErr(err)
	is.Equal(len(products), 0)

	// Range bounds become version suggestions, in version order.
	versions, err := st.Versions(ctx, "iterm2", "iterm2", "")
	is.NoErr(err)
	is.Equal(versions, []string{"3.5.6", "3.5.11"})

	is.NoErr(st.Optimize(ctx))
}
