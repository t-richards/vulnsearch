package serve

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// errNoTemplate is returned when rendering a page which doesn't exist.
var errNoTemplate = errors.New("no such template")

// funcs are the functions available to every template.
var funcs = template.FuncMap{
	"score": formatScore,
}

// parseTemplates parses each page in views/ together with the shared layouts
// and partials, keyed by the page's basename (e.g. "index"). Parsing happens
// once at startup, so template errors surface immediately.
func parseTemplates(files fs.FS) (map[string]*template.Template, error) {
	base, err := template.New("").Funcs(funcs).ParseFS(files, "views/layouts/*.html", "views/partials/*.html")
	if err != nil {
		return nil, fmt.Errorf("parsing layouts: %w", err)
	}

	pages, err := fs.Glob(files, "views/*.html")
	if err != nil {
		return nil, fmt.Errorf("listing pages: %w", err)
	}

	templates := map[string]*template.Template{}

	for _, page := range pages {
		tmpl, err := base.Clone()
		if err != nil {
			return nil, fmt.Errorf("cloning layouts: %w", err)
		}

		tmpl, err = tmpl.ParseFS(files, page)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", page, err)
		}

		templates[strings.TrimSuffix(path.Base(page), ".html")] = tmpl
	}

	return templates, nil
}

// render executes the named page into a buffer before writing it, so that a
// template error results in a clean 500 rather than a half-written page.
func (app *Application) render(w http.ResponseWriter, status int, page string, data any) {
	tmpl, ok := app.templates[page]
	if !ok {
		serverError(w, fmt.Errorf("%w: %s", errNoTemplate, page))

		return
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, "base", data)
	if err != nil {
		serverError(w, fmt.Errorf("rendering %s: %w", page, err))

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// formatScore formats a CVSS score, which may be missing.
func formatScore(score sql.NullFloat64) string {
	if !score.Valid {
		return "—"
	}

	return strconv.FormatFloat(score.Float64, 'f', 1, 64)
}
