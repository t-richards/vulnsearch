package cache

import (
	"log"
	"os"
	"path/filepath"
)

var cacheDir string

func init() {
	dir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalf("Could not determine cache directory: %v", err)
	}

	cacheDir = filepath.Join(dir, "vulnsearch")
	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		log.Fatalf("Could not create cache directory: %v", err)
	}
}

// DbPath gives the file path to the database
func DbPath(dbname string) string {
	override, ok := os.LookupEnv("VULNSEARCH_DB_PATH")
	if ok {
		return override
	}

	dbDir := filepath.Join(cacheDir, "db")
	err := os.MkdirAll(dbDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create db directory '%v': %v", dbDir, err)
	}
	return filepath.Join(dbDir, dbname)
}

// DataPath gives the directory path to the place where archives are cached
func DataPath() string {
	override, ok := os.LookupEnv("VULNSEARCH_DATA_PATH")
	if ok {
		return override
	}

	dataDir := filepath.Join(cacheDir, "data")
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create data directory '%v': %v", dataDir, err)
	}
	return dataDir
}
