package cache

import (
	"log"
	"os"
	"path/filepath"
)

var CacheDir string

func init() {
	dir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalf("Could not determine cache directory: %v", err)
	}

	CacheDir = filepath.Join(dir, "vulnsearch")
	err = os.MkdirAll(CacheDir, 0755)
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

	dbDir := filepath.Join(CacheDir, "db")
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

	dataDir := filepath.Join(CacheDir, "data")
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create data directory '%v': %v", dataDir, err)
	}
	return dataDir
}

// Clobber removes lots of data
func Clobber() {
	os.RemoveAll(CacheDir)
}
