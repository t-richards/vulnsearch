// Package serve provides the HTTP server and routes.
package serve

import (
	"compress/gzip"
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/t-richards/vulnsearch/internal/store"
)

// Static files.
//
//go:embed public
var publicFiles embed.FS

// HTML templates.
//
//go:embed views
var viewFiles embed.FS

// Application is an instance of the application with its dependencies.
type Application struct {
	http.Handler

	store     *store.Store
	templates map[string]*template.Template
	script    script
}

// NewApplication creates an application serving data from st.
func NewApplication(st *store.Store) (*Application, error) {
	templates, err := parseTemplates(viewFiles)
	if err != nil {
		return nil, err
	}

	script, err := buildScript()
	if err != nil {
		return nil, err
	}

	app := &Application{
		store:     st,
		templates: templates,
		script:    script,
	}

	public, err := fs.Sub(publicFiles, "public")
	if err != nil {
		return nil, fmt.Errorf("while loading public files: %w", err)
	}

	app.Handler = app.routes(public)

	return app, nil
}

// ListenAndServe serves the application on addr until ctx is canceled, then
// shuts down gracefully.
func (app *Application) ListenAndServe(ctx context.Context, addr string) error {
	server := &http.Server{
		Addr:              addr,
		Handler:           app,
		ReadHeaderTimeout: 10 * time.Second,
	}

	errs := make(chan error, 1)

	go func() {
		errs <- server.ListenAndServe()
	}()

	select {
	case err := <-errs:
		return fmt.Errorf("ListenAndServe: %w", err)
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
	defer cancel()

	err := server.Shutdown(shutdownCtx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("while shutting down: %w", err)
	}

	return nil
}

// routes creates a chi router (or "mux") with the application's routes and middleware.
func (app *Application) routes(public fs.FS) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.GetHead)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Compress(gzip.DefaultCompression))
	mux.Use(securityHeaders)

	// Pages.
	mux.Get("/", app.index)
	mux.Get("/search", app.search)

	// Search suggestions.
	mux.Get("/api/vendors", app.vendors)
	mux.Get("/api/products", app.products)
	mux.Get("/api/versions", app.versions)

	// Assets.
	mux.Get("/application.js", app.serveScript)
	mux.Get("/*", http.FileServerFS(public).ServeHTTP)

	return mux
}
