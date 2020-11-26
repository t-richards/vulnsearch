package nvd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

// Meta represents the meta status of a given year's feed
type Meta struct {
	LastModified time.Time
	Size         int
	ZipSize      int
	GzipSize     int
	Sha256       string
}

// DownloadMeta fetches the latest meta file from NVD and caches it
func DownloadMeta(year int) (*Meta, error) {
	url := metaURL(year)
	response, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch NVD meta status: %v", err)
	}
	defer response.Body.Close()

	meta, err := DecodeMeta(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse NVD metadata: %v", err)
	}

	return meta, nil
}

// DecodeMeta takes data from the reader and returns a meta object
func DecodeMeta(r io.Reader) (*Meta, error) {
	meta := new(Meta)

	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalf("Failed to read meta for unmarshalling")
	}

	lines := strings.Split(string(body), "\r\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		switch parts[0] {
		case "lastModifiedDate":
			meta.LastModified, _ = time.Parse(time.RFC3339, parts[1])
		case "size":
			meta.Size, _ = strconv.Atoi(parts[1])
		case "zipSize":
			meta.ZipSize, _ = strconv.Atoi(parts[1])
		case "gzSize":
			meta.GzipSize, _ = strconv.Atoi(parts[1])
		case "sha256":
			meta.Sha256 = parts[1]
		}
	}

	return meta, nil
}

func metaURL(year int) string {
	return fmt.Sprintf("%v/feeds/json/cve/%v/nvdcve-%v-%v.meta", BaseURI, Version, Version, year)
}
