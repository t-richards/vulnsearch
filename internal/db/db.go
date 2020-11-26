package db

import (
	"flag"
)

// Database constants
const (
	MainDb = "vulnsearch.sqlite3"
	TestDb = "vulnsearch_test.sqlite3"

	Schema = `
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

// Which tells you which database to use :)
func Which() string {
	if flag.Lookup("test.v") != nil {
		return TestDb
	}

	return MainDb
}
