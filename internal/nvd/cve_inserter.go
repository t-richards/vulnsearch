package nvd

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/t-richards/vulnsearch/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// LoadCVEs extracts compressed JSON archives of CVEs loads them into the DB
func LoadCVEs(db *gorm.DB) {
	currentYear := time.Now().Year()
	for year := EarliestYear; year <= currentYear; year++ {
		LoadCVEsForYear(db, year)
	}
}

// LoadCVEsForYear loads an already-downloaded archive
func LoadCVEsForYear(db *gorm.DB, year int) {
	path := ArchivePath(year)
	archive := loadFile(path)
	upsert(db, year, archive)
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

func upsert(db *gorm.DB, year int, archive Archive) {
	log.Printf("Processing %v archive...", year)

	for _, item := range archive.CveItems {
		// Create CVE entry
		cve := newCveFromCveItem(item)
		db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&cve)

		// Create relationships between the things
		relatedProducts := findRelatedProductsForItem(db, item)
		db.Clauses(clause.OnConflict{UpdateAll: true}).
			Model(&cve).Association("Products").Append(relatedProducts)
	}
}

func findRelatedProductsForItem(db *gorm.DB, item CveItem) []models.Product {
	result := make([]models.Product, 0)

	for _, node := range item.Configurations.Nodes {
		if node.Operator == "OR" {
			for _, match := range node.CpeMatch {
				product := parseProduct(match.Cpe23Url)
				db.Where(&product).FirstOrInit(&product)
				result = append(result, product)
			}
		}
	}

	return result
}

func parseProduct(cpe23url string) models.Product {
	parts := strings.Split(cpe23url, ":")
	if len(parts) < 6 {
		log.Fatalf("Unknown cpe23url format: %v", cpe23url)
	}

	return models.Product{
		Vendor:  parts[3],
		Name:    parts[4],
		Version: parts[5],
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
