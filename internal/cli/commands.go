package cli

import (
	"log"
	"vulnsearch/internal/app"
	"vulnsearch/internal/db"
)

// Migrate migrates the database
func Migrate() {
	app := app.New()
	app.DB.Exec(db.Schema)
}

// Optimize cleans up the database after an import
func Optimize() {
	log.Printf("Optimizing database...")
	app := app.New()
	app.DB.Exec("PRAGMA synchronous = OFF")
	app.DB.Exec("PRAGMA journal_mode = memory")
	app.DB.Exec("PRAGMA optimize")
	app.DB.Exec("VACUUM")
	log.Printf("Done.")
}
