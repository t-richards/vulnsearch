package nvd

import (
	"compress/gzip"
	"encoding/xml"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/t-richards/vulnsearch/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// LoadProducts loads products into the database
func LoadProducts(db *gorm.DB) {
	log.Printf("[BEGIN] Loading products")

	cpeList := decodeProductsArchive()
	for _, item := range cpeList.Items {
		product := productFromCpeItem(item)
		db.Clauses(clause.OnConflict{DoNothing: true}).Create(&product)
	}

	log.Printf("[END] Loading products")
}

func decodeProductsArchive() CpeList {
	archivePath := ProductsPath()

	file, err := os.Open(archivePath)
	if err != nil {
		log.Fatalf("Failed to open products archive: %v", err)
	}

	reader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatalf("Failed to decompress products archive: %v", err)
	}

	cpeList := CpeList{}
	err = xml.NewDecoder(reader).Decode(&cpeList)
	if err != nil {
		log.Fatalf("Failed to decode products archive XML: %v", err)
	}

	return cpeList
}

func productFromCpeItem(item CpeItem) models.Product {
	parts := strings.Split(item.Name, ":")
	if len(parts) < 5 {
		log.Fatalf("Invalid parts for item: %v", item.Name)
	}

	return models.Product{
		Vendor:  unescapePart(parts[2]),
		Name:    unescapePart(parts[3]),
		Version: unescapePart(parts[4]),
	}
}

// It preferable to persist unescaped values, but it isn't critical
func unescapePart(input string) string {
	result, err := url.QueryUnescape(input)
	if err != nil {
		return input
	}

	return result
}
