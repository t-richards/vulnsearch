package db

import (
	_ "embed"
	"flag"
	"log"

	"github.com/t-richards/vulnsearch/internal/cache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed schema.sql
var schema string

// Database constants
const (
	mainDB = "vulnsearch.sqlite3"
	testDB = "vulnsearch_test.sqlite3"

	enableForeignKeys = `
	PRAGMA foreign_keys = ON;
	`

	fastJournal = `
	PRAGMA synchronous = OFF;
	PRAGMA journal_mode = memory;
	`

	optimize = `
	PRAGMA optimize;
	VACUUM;
	`
)

// Which tells you the basename of the database file
func Which() string {
	if flag.Lookup("test.v") != nil {
		return testDB
	}

	return mainDB
}

// FastJournal sacrifices data reliability for insert speed
func FastJournal(conn *gorm.DB) {
	conn.Exec(fastJournal)
}

// Migrate migrates the database
func Migrate(conn *gorm.DB) {
	conn.Exec(schema)
}

func Connect() *gorm.DB {
	dbPath := cache.DbPath(Which())
	conn, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database '%v': %v", dbPath, err)
	}
	conn.Exec(enableForeignKeys)

	return conn
}

// Optimize optimizes the database
func Optimize(conn *gorm.DB) {
	conn.Exec(fastJournal)
	conn.Exec(optimize)
}
