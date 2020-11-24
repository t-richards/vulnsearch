package commands

import (
	"log"
	"time"
	"vulnsearch/internal/nvd"
)

// Fetch downloads missing and updated data files for all years
func Fetch() {
	currentYear := time.Now().Year()

	for year := 2002; year <= currentYear; year++ {
		if nvd.NeedsDownload(year) {
			log.Printf("Downloading: %v\n", year)
			nvd.Download(year)
		} else {
			log.Printf("Already downloaded: %v\n", year)
		}
	}
}
