package app

import (
	"encoding/json"
	"log"
	"net/http"
	"vulnsearch/internal/db"
	"vulnsearch/internal/views"
	"vulnsearch/internal/webserver"

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
	app.DB, err = gorm.Open(sqlite.Open(db.Which()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Router
	app.Router = httprouter.New()
	app.setupRoutes()

	return app
}

func (app *App) setupRoutes() {
	app.Router.GET("/", app.index())
	app.Router.POST("/vendor", app.vendor())
	app.Router.NotFound = webserver.PublicFiles
}

// FastMode sacrifices reliability for performance
func (app *App) FastMode() {
	app.DB.Exec("PRAGMA synchronous = OFF")
	app.DB.Exec("PRAGMA journal_mode = memory")
}

// Run starts the server
func (app *App) Run() {
	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

func (app *App) index() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		views.HomeTemplate.Execute(w, nil)
	}
}

func (app *App) vendor() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		params := SearchParams{}
		json.NewDecoder(r.Body).Decode(&params)
		defer r.Body.Close()

		resp := VendorResponse{
			Vendors: make([]string, 0),
		}
		like := "%" + *params.Vendor + "%"
		app.DB.Table("products").Select("DISTINCT vendor").Where("vendor LIKE ?", like).Order("vendor ASC").Limit(100).Find(&resp.Vendors)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resp)
	}
}
