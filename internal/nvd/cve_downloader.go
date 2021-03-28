package nvd

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/t-richards/vulnsearch/internal/cache"
)

func downloadCVEs() {
	log.Printf("[BEGIN] Download CVEs")
	currentYear := time.Now().Year()

	for year := EarliestYear; year <= currentYear; year++ {
		if cveNeedsDownload(year) {
			log.Printf("Downloading CVEs for: %v\n", year)
			sourceURL := archiveURL(year)
			destPath := ArchivePath(year)
			download(sourceURL, destPath)
		} else {
			log.Printf("Already downloaded CVEs for: %v\n", year)
		}
	}
	log.Printf("[END] Download CVEs")
}

func cveNeedsDownload(year int) bool {
	targetPath := ArchivePath(year)

	// The data file should be downloaded if it doesn't exist
	stat, err := os.Stat(targetPath)
	if err != nil {
		return true
	}

	// Fetch metadata about the given year's archive
	meta, err := DownloadMeta(year)
	if err != nil {
		return true
	}

	// The archive should be re-downloaded if it is out of date
	if meta.LastModified.After(stat.ModTime()) {
		return true
	}

	return false
}

// ArchivePath returns the path to the gzipped JSON archive for a given year
func ArchivePath(year int) string {
	basename := fmt.Sprintf("nvdcve-%v-%v.json.gz", Version, year)
	return path.Join(cache.DataPath(), basename)
}

func archiveURL(year int) string {
	return fmt.Sprintf("%v/feeds/json/cve/%v/nvdcve-%v-%v.json.gz", BaseURI, Version, Version, year)
}
