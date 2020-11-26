package nvd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/t-richards/vulnsearch/internal/cache"
)

// Products constants
const (
	ProductsBasename = "official-cpe-dictionary_v2.3.xml.gz"
)

func downloadProducts() {
	log.Printf("[BEGIN] Download Products")
	sourceURL := productsURL()
	destPath := ProductsPath()
	download(sourceURL, destPath)
	log.Printf("[END] Download Products")
}

func productsURL() string {
	return fmt.Sprintf("%v/feeds/xml/cpe/dictionary/%v", BaseURI, ProductsBasename)
}

// ProductsPath is the path of the products archive on disk
func ProductsPath() string {
	return filepath.Join(cache.DataPath(), ProductsBasename)
}
