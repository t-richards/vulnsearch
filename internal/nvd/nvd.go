package nvd

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Package-level constants
const (
	BaseURI      = "https://nvd.nist.gov"
	EarliestYear = 2002
	Version      = "1.1"
)

var client = http.Client{
	Timeout: 60 * time.Second,
}

// DownloadAll fetches JSON archives for all years
func DownloadAll() {
	downloadCVEs()
}

func download(sourceURL string, destPath string) {
	response, err := client.Get(sourceURL)
	if err != nil {
		log.Printf("Failed to download '%v': %v", sourceURL, err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(destPath)
	if err != nil {
		log.Printf("Failed to create output file %v: %v", destPath, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Printf("Failed to write archive file data: %v", err)
	}
}
