package app

import (
	"log"
	"net/http"

	"github.com/t-richards/vulnsearch/internal/cache"
	"github.com/t-richards/vulnsearch/internal/db"
	"github.com/t-richards/vulnsearch/internal/webserver"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App is a small combination of a database connection + http router
type App struct {
	Router *httprouter.Router
	DB     *gorm.DB
}

// New returns an initialized app instance
func New() *App {
	app := new(App)

	// Database
	var err error
	dbPath := cache.DbPath(db.Which())
	app.DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database '%v': %v", dbPath, err)
	}

	// Router
	app.Router = httprouter.New()
	app.setupRoutes()

	return app
}

func (app *App) setupRoutes() {
	app.Router.GET("/", app.index())
	app.Router.POST("/vendor", app.vendor())
	app.Router.POST("/product", app.product())
	app.Router.POST("/version", app.version())
	app.Router.NotFound = webserver.PublicFiles
}

// Run starts the HTTP server
func (app *App) Run() {
	log.Println("Listening on http://localhost:5000/")
	log.Fatal(http.ListenAndServe(":5000", app.Router))
}
