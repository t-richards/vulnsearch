package nvd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// Meta represents the meta status of a given year's feed.
type Meta struct {
	LastModified time.Time
	Size         int
	ZipSize      int
	GzipSize     int
	Sha256       string
}

// FetchMeta fetches the meta file for the given year's feed.
func FetchMeta(ctx context.Context, year int) (Meta, error) {
	resp, err := get(ctx, metaURL(year))
	if err != nil {
		return Meta{}, err
	}
	defer resp.Body.Close()

	return DecodeMeta(resp.Body)
}

// DecodeMeta parses a meta file, which consists of "key:value" lines.
func DecodeMeta(r io.Reader) (Meta, error) {
	var meta Meta

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		key, value, ok := strings.Cut(strings.TrimSpace(scanner.Text()), ":")
		if !ok {
			continue
		}

		var err error

		switch key {
		case "lastModifiedDate":
			meta.LastModified, err = time.Parse(time.RFC3339, value)
		case "size":
			meta.Size, err = strconv.Atoi(value)
		case "zipSize":
			meta.ZipSize, err = strconv.Atoi(value)
		case "gzSize":
			meta.GzipSize, err = strconv.Atoi(value)
		case "sha256":
			meta.Sha256 = value
		}

		if err != nil {
			return Meta{}, fmt.Errorf("while parsing meta %s: %w", key, err)
		}
	}

	err := scanner.Err()
	if err != nil {
		return Meta{}, fmt.Errorf("while reading meta: %w", err)
	}

	return meta, nil
}
