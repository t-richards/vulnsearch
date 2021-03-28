package app

import (
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
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
	engine := html.NewFileSystem(http.FS(views.Views), ".html")
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
	app.Get("/search", search)

	// Assets
	assetDir, _ := fs.Sub(assets.Assets, "public")
	app.Use("/", cache.New(cache.Config{
		CacheControl: true,
	}))
	app.Use("/", compress.New())
	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(assetDir),
		Browse: true,
	}))

	// Start!
	app.Listen(":3000")
}
