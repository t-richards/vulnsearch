// Package store opens the SQLite database, applies the schema on startup, and
// exposes the sqlc queries alongside higher-level lookups and imports.
package store

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "modernc.org/sqlite" // Registers the "sqlite" driver.

	"github.com/t-richards/vulnsearch/internal/db"
	"github.com/t-richards/vulnsearch/internal/sqlc"
)

// Store wraps the generated queries and the underlying pool.
type Store struct {
	*db.Queries

	conn *sql.DB
}

// Open opens (creating if needed) the database at path and applies the schema.
func Open(ctx context.Context, path string) (*Store, error) {
	// Pragmas in the DSN apply to every connection in the pool, rather than
	// only whichever connection happens to run a one-off PRAGMA statement.
	params := url.Values{}
	params.Add("_pragma", "foreign_keys(1)")
	params.Add("_pragma", "journal_mode(WAL)")
	params.Add("_pragma", "synchronous(NORMAL)")
	params.Add("_pragma", "busy_timeout(5000)")
	params.Add("_time_format", "sqlite")

	conn, err := sql.Open("sqlite", "file:"+path+"?"+params.Encode())
	if err != nil {
		return nil, fmt.Errorf("while opening database: %w", err)
	}

	_, err = conn.ExecContext(ctx, sqlc.Schema)
	if err != nil {
		conn.Close()

		return nil, fmt.Errorf("while applying schema: %w", err)
	}

	return &Store{Queries: db.New(conn), conn: conn}, nil
}

// Begin starts a transaction and returns it alongside queries bound to it.
// The caller commits; defer tx.Rollback().
func (s *Store) Begin(ctx context.Context) (*sql.Tx, *db.Queries, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("while beginning transaction: %w", err)
	}

	return tx, s.WithTx(tx), nil
}

// Close closes the pool.
func (s *Store) Close() error {
	err := s.conn.Close()
	if err != nil {
		return fmt.Errorf("while closing database: %w", err)
	}

	return nil
}

// Optimize updates query planner statistics and reclaims unused space.
func (s *Store) Optimize(ctx context.Context) error {
	_, err := s.conn.ExecContext(ctx, "PRAGMA optimize; VACUUM;")
	if err != nil {
		return fmt.Errorf("while optimizing database: %w", err)
	}

	return nil
}
