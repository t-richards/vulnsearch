package nvd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var client http.Client

func init() {
	client = http.Client{
		Timeout: 60 * time.Second,
	}
}

// DownloadAll fetches JSON archives for all years
func DownloadAll() {
	currentYear := time.Now().Year()

	for year := EarliestYear; year <= currentYear; year++ {
		if needsDownload(year) {
			log.Printf("Downloading: %v\n", year)
			download(year)
		} else {
			log.Printf("Already downloaded: %v\n", year)
		}
	}
}

func download(year int) {
	outFilePath := ArchivePath(year)

	url := archiveURL(year)
	response, err := client.Get(url)
	if err != nil {
		log.Printf("Failed to download year %v: %v", year, err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(outFilePath)
	if err != nil {
		log.Printf("Failed to create output file %v: %v", outFilePath, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Printf("Failed to write archive file data: %v", err)
	}
}

func needsDownload(year int) bool {
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

// ArchivePath returns a relative path to the gzipped JSON archive for a given year
func ArchivePath(year int) string {
	basename := fmt.Sprintf("nvdcve-%v-%v.json.gz", Version, year)
	return path.Join(DataDir, basename)
}

func archiveURL(year int) string {
	return fmt.Sprintf("%v/feeds/json/cve/%v/nvdcve-%v-%v.json.gz", BaseURI, Version, Version, year)
}
