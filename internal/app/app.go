package app

import (
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
	"github.com/t-richards/vulnsearch/internal/assets"
	"github.com/t-richards/vulnsearch/internal/db"
	"github.com/t-richards/vulnsearch/internal/views"
	"gorm.io/gorm"
)

var conn *gorm.DB

// Run starts the application
func Run() {
	// Database
	conn = db.Connect()

	// Router
	templateDir, _ := fs.Sub(views.Templates, "templates")
	engine := html.NewFileSystem(http.FS(templateDir), ".html")
	engine.Layout("content")
	config := fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Vulnsearch",
		Views:         engine,
	}
	app := fiber.New(config)
	app.Use(recover.New())
	app.Use(requestid.New())

	// Routes
	app.Get("/", index)
	app.Get("/application.js", build)
	app.Post("/vendor", vendor)
	app.Post("/product", product)
	app.Post("/version", version)
	app.Post("/search", search)

	// Assets
	assetDir, _ := fs.Sub(assets.Assets, "public")
	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(assetDir),
		Browse: true,
	}))

	// Start!
	app.Listen(":3000")
}
