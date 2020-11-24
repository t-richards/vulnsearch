package nvd

import (
	"fmt"
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

// Download fetches a JSON file for a given year
func Download(year int) {
	// targetPath := dataFilePath(year)

}

// NeedsDownload determines whether a particular year's archive needs downloading
func NeedsDownload(year int) bool {
	targetPath := dataFilePath(year)

	_, err := os.Stat(targetPath)
	if err != nil {
		return true
	}

	return false
}

func dataFilePath(year int) string {
	basename := fmt.Sprintf("nvdcve-%v-%v.json.gz", Version, year)
	return path.Join(DataDir, basename)
}
