// Package paths locates the database and downloaded data feeds on disk.
//
// By default, all files live under os.UserCacheDir(), since everything can be
// re-downloaded and rebuilt from the NVD feeds.
package paths

import (
	"fmt"
	"os"
	"path/filepath"
)

// Environment variables which override the default locations.
const (
	DBPathEnv   = "VULNSEARCH_DB_PATH"
	DataPathEnv = "VULNSEARCH_DATA_PATH"
)

// CacheDir returns the application's cache directory.
func CacheDir() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("finding cache directory: %w", err)
	}

	return filepath.Join(dir, "vulnsearch"), nil
}

// DBPath returns the path to the SQLite database, creating its parent
// directory if needed.
func DBPath() (string, error) {
	path, ok := os.LookupEnv(DBPathEnv)
	if !ok {
		dir, err := CacheDir()
		if err != nil {
			return "", err
		}

		path = filepath.Join(dir, "db", "vulnsearch.sqlite3")
	}

	err := os.MkdirAll(filepath.Dir(path), 0o750)
	if err != nil {
		return "", fmt.Errorf("creating database directory: %w", err)
	}

	return path, nil
}

// DataDir returns the directory containing downloaded feeds, creating it if needed.
func DataDir() (string, error) {
	dir, ok := os.LookupEnv(DataPathEnv)
	if !ok {
		cache, err := CacheDir()
		if err != nil {
			return "", err
		}

		dir = filepath.Join(cache, "data")
	}

	err := os.MkdirAll(dir, 0o750)
	if err != nil {
		return "", fmt.Errorf("creating data directory: %w", err)
	}

	return dir, nil
}
