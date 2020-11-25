package nvd

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// LoadFiles extracts compressed JSON archives and loads them into the DB
func LoadFiles() {
	currentYear := time.Now().Year()

	for year := EarliestYear; year <= currentYear; year++ {
		path := ArchivePath(year)
		archive := loadFile(path)
		upsert(year, archive)
	}
}

func loadFile(path string) Archive {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open archive %v: %v", path, err)
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatalf("Failed to decompress archive %v: %v", path, err)
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("Failed to read archive %v: %v", path, err)
	}

	archive := Archive{}
	err = json.Unmarshal(data, &archive)
	if err != nil {
		log.Fatalf("Failed to parse JSON in archive %v: %v", path, err)
	}

	return archive
}

func upsert(year int, archive Archive) {
	log.Printf("Processing %v archive...", year)

	max := 0

	for i := range archive.CveItems {
		max = i
	}

	log.Printf("%v items loaded!", max)
}
