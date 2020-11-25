package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connection *gorm.DB

// Database constants
const (
	devDb  = "db/vulnsearch_dev.sqlite3"
	testDb = "db/vulnsearch_test.sqlite3"
	schema = `
	CREATE TABLE IF NOT EXISTS cves (
		id            VARCHAR(64) NOT NULL PRIMARY KEY,
		description   TEXT        NOT NULL,
		cwe_id        VARCHAR(32) NOT NULL,
		cvss_v2_score REAL        NOT NULL,
		cvss_v3_score REAL        NOT NULL,
		published     DATETIME    NOT NULL,
		last_modified DATETIME    NOT NULL
	);

	CREATE TABLE IF NOT EXISTS products (
		id      INTEGER      NOT NULL PRIMARY KEY,
		vendor  VARCHAR(255) NOT NULL,
		name    VARCHAR(255) NOT NULL,
		version VARCHAR(255) NOT NULL
	);

	CREATE UNIQUE INDEX IF NOT EXISTS products_unique
		ON products (vendor, name, version);

	CREATE TABLE IF NOT EXISTS cves_products (
		cve_id     VARCHAR(64) NOT NULL,
		product_id INTEGER     NOT NULL
	);

	CREATE INDEX IF NOT EXISTS cves_products_cve_id
		ON cves_products (cve_id);
	CREATE INDEX IF NOT EXISTS cves_products_product_id
		ON cves_products (product_id);

	CREATE UNIQUE INDEX IF NOT EXISTS cves_products_unique
		ON cves_products (cve_id, product_id);
	`
)

func init() {
	var err error
	connection, err = gorm.Open(sqlite.Open(devDb), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

// Migrate attempts to apply the schema to the database
func Migrate() {
	connection.Exec(schema)
}
