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

func DbPath(dbname string) string {
	override, ok := os.LookupEnv("VULNSEARCH_DB_PATH")
	if ok {
		return override
	}

	dbDir := filepath.Join(cacheDir, "db")
	os.MkdirAll(dbDir, 0755)
	return filepath.Join(dbDir, dbname)
}

func DataPath() string {
	override, ok := os.LookupEnv("VULNSEARCH_DATA_PATH")
	if ok {
		return override
	}

	dataDir := filepath.Join(cacheDir, "data")
	os.MkdirAll(dataDir, 0755)
	return dataDir
}
