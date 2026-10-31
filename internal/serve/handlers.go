package serve

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/t-richards/vulnsearch/internal/store"
)

func (app *Application) index(w http.ResponseWriter, _ *http.Request) {
	app.render(w, http.StatusOK, "index", nil)
}

// productPage is the data for the product template.
type productPage struct {
	Vendor     string
	Product    string
	Version    string
	CVEs       []store.AffectedCVE
	Severities []severityTable
}

// FullName is a human-readable name for the product version.
func (p productPage) FullName() string {
	return strings.TrimSpace(strings.Join([]string{p.Vendor, p.Product, p.Version}, " "))
}

func (app *Application) search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := productPage{
		Vendor:  strings.TrimSpace(query.Get("vendor")),
		Product: strings.TrimSpace(query.Get("name")),
		Version: strings.TrimSpace(query.Get("version")),
	}

	if page.Vendor == "" || page.Product == "" {
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	cves, err := app.store.AffectedCVEs(r.Context(), page.Vendor, page.Product, page.Version)
	if err != nil {
		serverError(w, err)

		return
	}

	page.CVEs = cves
	page.Severities = severityTables(cves)

	app.render(w, http.StatusOK, "product", page)
}

func (app *Application) vendors(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	vendors, err := app.store.Vendors(r.Context(), query.Get("vendor"))
	if err != nil {
		serverError(w, err)

		return
	}

	writeJSON(w, vendors)
}

func (app *Application) products(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	products, err := app.store.Products(r.Context(), query.Get("vendor"), query.Get("name"))
	if err != nil {
		serverError(w, err)

		return
	}

	writeJSON(w, products)
}

func (app *Application) versions(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	versions, err := app.store.Versions(r.Context(), query.Get("vendor"), query.Get("name"), query.Get("version"))
	if err != nil {
		serverError(w, err)

		return
	}

	writeJSON(w, versions)
}

// writeJSON writes a list of strings as a JSON array (never null).
func writeJSON(w http.ResponseWriter, list []string) {
	if list == nil {
		list = []string{}
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Println(err) // The client likely went away; headers are already sent.
	}
}
