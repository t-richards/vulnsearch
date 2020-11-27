package db

import (
	"flag"

	"gorm.io/gorm"
)

// Database constants
const (
	mainDB = "vulnsearch.sqlite3"
	testDB = "vulnsearch_test.sqlite3"

	schema = `
	/* CVEs table */
	CREATE TABLE IF NOT EXISTS cves (
		id            TEXT     NOT NULL PRIMARY KEY,
		description   TEXT     NOT NULL,
		cwe_id        TEXT     NOT NULL,
		cvss_v2_score REAL     NOT NULL,
		cvss_v3_score REAL     NOT NULL,
		published     DATETIME NOT NULL,
		last_modified DATETIME NOT NULL
	);

	/* Products table */
	CREATE TABLE IF NOT EXISTS products (
		id      INTEGER NOT NULL PRIMARY KEY,
		vendor  TEXT    NOT NULL,
		name    TEXT    NOT NULL,
		version TEXT    NOT NULL
	);

	CREATE UNIQUE INDEX IF NOT EXISTS products_unique
		ON products (vendor, name, version);

	/* Search indexes */
	CREATE INDEX IF NOT EXISTS products_vendor
		ON products (vendor);
	CREATE INDEX IF NOT EXISTS products_vendor_name
		ON products (vendor, name);

	/* CVEs <-> Products */
	CREATE TABLE IF NOT EXISTS cves_products (
		cve_id     TEXT    NOT NULL,
		product_id INTEGER NOT NULL,

		PRIMARY KEY (cve_id, product_id)
	);

	CREATE INDEX IF NOT EXISTS cves_products_cve_id
		ON cves_products (cve_id);
	CREATE INDEX IF NOT EXISTS cves_products_product_id
		ON cves_products (product_id);
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

// Migrate migrates the database
func Migrate(db *gorm.DB) {
	db.Exec(schema)
}

// FastJournal sacrifices data reliability for insert speed
func FastJournal(db *gorm.DB) {
	db.Exec(fastJournal)
}

// Optimize optimizes the database
func Optimize(db *gorm.DB) {
	db.Exec(fastJournal)
	db.Exec(optimize)
}
