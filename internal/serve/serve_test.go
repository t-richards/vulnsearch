package serve_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matryer/is"

	"github.com/t-richards/vulnsearch/internal/nvd"
	"github.com/t-richards/vulnsearch/internal/serve"
	"github.com/t-richards/vulnsearch/internal/store"
)

// newApp creates an application backed by a database loaded with the nvd
// package's test feed.
func newApp(t *testing.T) *serve.Application {
	t.Helper()

	is := is.New(t)

	st, err := store.Open(t.Context(), filepath.Join(t.TempDir(), "test.sqlite3"))
	is.NoErr(err)
	t.Cleanup(func() { st.Close() })

	file, err := os.Open("../nvd/testdata/feed.json")
	is.NoErr(err)

	defer file.Close()

	_, err = st.Import(t.Context(), nvd.Decode(file))
	is.NoErr(err)
	is.NoErr(st.RebuildProducts(t.Context()))

	app, err := serve.NewApplication(st)
	is.NoErr(err)

	return app
}

// response is the result of a test request.
type response struct {
	StatusCode int
	Header     http.Header
	Body       string
}

// get performs a GET request against app.
func get(t *testing.T, app http.Handler, target string) response {
	t.Helper()

	is := is.New(t)

	r := httptest.NewRequestWithContext(t.Context(), http.MethodGet, target, nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	is.NoErr(err)

	return response{StatusCode: res.StatusCode, Header: res.Header, Body: string(body)}
}

func TestIndex(t *testing.T) {
	is := is.New(t)
	app := newApp(t)

	res := get(t, app, "/")

	is.Equal(res.StatusCode, http.StatusOK)
	is.Equal(res.Header.Get("Content-Type"), "text/html; charset=utf-8")
	is.True(strings.Contains(res.Body, "Start typing to find a product"))
	is.True(strings.Contains(res.Body, "bootstrap.min.css"))
	is.True(res.Header.Get("Content-Security-Policy") != "")
}

func TestSearchAffected(t *testing.T) {
	is := is.New(t)
	app := newApp(t)

	res := get(t, app, "/search?vendor=iterm2&name=iterm2&version=3.5.10")

	is.Equal(res.StatusCode, http.StatusOK)
	is.True(strings.Contains(res.Body, "<title>iterm2 iterm2 3.5.10 - Vulnsearch</title>"))
	is.True(strings.Contains(res.Body, "CVE-2025-22275"))
	is.True(strings.Contains(res.Body, "&gt;= 3.5.6, &lt; 3.5.11"))
	is.True(strings.Contains(res.Body, "9.3"))
}

func TestSearchUnaffected(t *testing.T) {
	is := is.New(t)
	app := newApp(t)

	res := get(t, app, "/search?vendor=iterm2&name=iterm2&version=3.5.11")

	is.Equal(res.StatusCode, http.StatusOK)
	is.True(!strings.Contains(res.Body, "CVE-2025-22275"))
	is.True(strings.Contains(res.Body, "No known CVEs"))
}

func TestSearchMissingParams(t *testing.T) {
	is := is.New(t)
	app := newApp(t)

	res := get(t, app, "/search?vendor=iterm2")

	is.Equal(res.StatusCode, http.StatusFound)
	is.Equal(res.Header.Get("Location"), "/")
}

func TestSuggestions(t *testing.T) {
	tests := []struct {
		target string
		want   []string
	}{
		{"/api/vendors?vendor=iter", []string{"iterm2"}},
		{"/api/vendors?vendor=nope", []string{}},
		{"/api/products?vendor=joomlabiblestudy&name=com", []string{"com_biblestudy"}},
		{"/api/versions?vendor=iterm2&name=iterm2&version=3.5", []string{"3.5.6", "3.5.11"}},
	}

	app := newApp(t)

	for _, tt := range tests {
		t.Run(tt.target, func(t *testing.T) {
			is := is.New(t)

			res := get(t, app, tt.target)
			is.Equal(res.StatusCode, http.StatusOK)
			is.Equal(res.Header.Get("Content-Type"), "application/json")

			var got []string

			is.NoErr(json.Unmarshal([]byte(res.Body), &got))
			is.Equal(got, tt.want)
		})
	}
}

func TestStaticFiles(t *testing.T) {
	is := is.New(t)
	app := newApp(t)

	res := get(t, app, "/application.js")
	is.Equal(res.StatusCode, http.StatusOK)
	is.True(strings.HasPrefix(res.Header.Get("Content-Type"), "text/javascript"))
	is.True(strings.Contains(res.Body, "/api/vendors"))

	res = get(t, app, "/favicon.ico")
	is.Equal(res.StatusCode, http.StatusOK)
}

func TestNotFound(t *testing.T) {
	is := is.New(t)
	app := newApp(t)

	res := get(t, app, "/not-found")
	is.Equal(res.StatusCode, http.StatusNotFound)
}

func TestHead(t *testing.T) {
	app := newApp(t)

	for _, target := range []string{"/", "/search?vendor=iterm2&name=iterm2", "/api/vendors", "/favicon.ico"} {
		t.Run(target, func(t *testing.T) {
			is := is.New(t)

			r := httptest.NewRequestWithContext(t.Context(), http.MethodHead, target, nil)
			w := httptest.NewRecorder()
			app.ServeHTTP(w, r)

			is.Equal(w.Code, http.StatusOK)
		})
	}
}

func TestScriptCaching(t *testing.T) {
	is := is.New(t)
	app := newApp(t)

	res := get(t, app, "/application.js")
	is.Equal(res.StatusCode, http.StatusOK)
	is.Equal(res.Header.Get("Cache-Control"), "no-cache")

	etag := res.Header.Get("ETag")
	is.True(etag != "")

	// A matching ETag is answered without a body.
	r := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/application.js", nil)
	r.Header.Set("If-None-Match", etag)

	w := httptest.NewRecorder()
	app.ServeHTTP(w, r)

	is.Equal(w.Code, http.StatusNotModified)
	is.Equal(w.Body.Len(), 0)
}
