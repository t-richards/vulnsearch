// Package sqlc holds the embedded database schema.
package sqlc

import _ "embed"

//go:generate sqlc generate

// Schema is the DDL applied to initialize the database on startup.
//
//go:embed schema.sql
var Schema string
