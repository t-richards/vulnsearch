// Package nvd downloads and decodes the NVD CVE JSON 2.0 data feeds.
//
// See https://nvd.nist.gov/vuln/data-feeds for details about the feeds.
package nvd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Feed constants.
const (
	BaseURL      = "https://nvd.nist.gov/feeds/json/cve/2.0"
	EarliestYear = 2002 // The 2002 feed also contains all CVEs from 1999-2001.
	userAgent    = "vulnsearch (+https://github.com/t-richards/vulnsearch)"
)

// errUnexpectedStatus is returned for non-200 responses from NVD.
var errUnexpectedStatus = errors.New("unexpected status")

var client = &http.Client{
	Timeout: 5 * time.Minute,
}

// Years returns every year which has a CVE feed.
func Years() []int {
	var years []int

	for year := EarliestYear; year <= time.Now().Year(); year++ {
		years = append(years, year)
	}

	return years
}

// ArchiveName is the basename of the gzipped JSON feed for a given year.
func ArchiveName(year int) string {
	return fmt.Sprintf("nvdcve-2.0-%d.json.gz", year)
}

func archiveURL(year int) string {
	return BaseURL + "/" + ArchiveName(year)
}

func metaURL(year int) string {
	return fmt.Sprintf("%s/nvdcve-2.0-%d.meta", BaseURL, year)
}

// Fetch downloads the feed for each year into dir, skipping any feeds which
// are already up to date.
func Fetch(ctx context.Context, dir string, years []int) error {
	for _, year := range years {
		dest := filepath.Join(dir, ArchiveName(year))

		stale, err := needsDownload(ctx, year, dest)
		if err != nil {
			return err
		}

		if !stale {
			log.Printf("Already up to date: %d", year)

			continue
		}

		log.Printf("Downloading: %d", year)

		err = download(ctx, archiveURL(year), dest)
		if err != nil {
			return err
		}
	}

	return nil
}

// needsDownload reports whether the archive at path is missing or older than the
// feed's most recent modification.
func needsDownload(ctx context.Context, year int, path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return true, nil //nolint:nilerr // A missing archive just needs downloading.
	}

	meta, err := FetchMeta(ctx, year)
	if err != nil {
		return false, err
	}

	return meta.LastModified.After(stat.ModTime()), nil
}

func get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()

		return nil, fmt.Errorf("fetching %s: %w %s", url, errUnexpectedStatus, resp.Status)
	}

	return resp, nil
}

// download writes url to dest via a temporary file, so an interrupted download
// never leaves a truncated archive behind.
func download(ctx context.Context, url, dest string) error {
	resp, err := get(ctx, url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmp, err := os.CreateTemp(filepath.Dir(dest), filepath.Base(dest)+".*.tmp")
	if err != nil {
		return fmt.Errorf("creating temporary file: %w", err)
	}
	defer os.Remove(tmp.Name())

	_, err = io.Copy(tmp, resp.Body)
	if err != nil {
		tmp.Close()

		return fmt.Errorf("downloading %s: %w", url, err)
	}

	err = tmp.Close()
	if err != nil {
		return fmt.Errorf("writing %s: %w", tmp.Name(), err)
	}

	err = os.Rename(tmp.Name(), dest)
	if err != nil {
		return fmt.Errorf("saving %s: %w", dest, err)
	}

	return nil
}
