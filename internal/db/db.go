package db

import (
	"flag"
)

// Database constants
const (
	MainDb = "vulnsearch.sqlite3"
	TestDb = "vulnsearch_test.sqlite3"

	Schema = `
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
		vendor  TEXT NOT NULL,
		name    TEXT NOT NULL,
		version TEXT NOT NULL,

		PRIMARY KEY(vendor, name, version)
	);

	/* Search indexes */
	CREATE INDEX IF NOT EXISTS products_vendor
		ON products (vendor);
	CREATE INDEX IF NOT EXISTS products_vendor_name
		ON products (vendor, name);
	`
)

// Which tells you the basename of the database file
func Which() string {
	if flag.Lookup("test.v") != nil {
		return TestDb
	}

	return MainDb
}
