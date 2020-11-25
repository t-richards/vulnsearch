package nvd

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
	"vulnsearch/internal/db"
	"vulnsearch/internal/models"

	"gorm.io/gorm/clause"
)

// LoadFiles extracts compressed JSON archives and loads them into the DB
func LoadFiles() {
	db.FastMode()
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
	conn := db.GetConnection()
	log.Printf("Processing %v archive...", year)

	for _, item := range archive.CveItems {
		cve := newCveFromCveItem(item)

		conn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&cve)
	}
}

func newCveFromCveItem(item CveItem) models.Cve {
	return models.Cve{
		ID:           item.Cve.CveMeta.ID,
		Description:  item.Description(),
		CweID:        item.CweID(),
		CvssV2Score:  item.Impact.BaseMetricV2.CvssV2.BaseScore,
		CvssV3Score:  item.Impact.BaseMetricV3.CvssV3.BaseScore,
		Published:    item.Published.Time,
		LastModified: item.LastModified.Time,
	}
}
