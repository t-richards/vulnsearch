package nvd_test

import (
	"os"
	"testing"
	"time"

	"github.com/matryer/is"

	"github.com/t-richards/vulnsearch/internal/nvd"
)

func TestDecodeMeta(t *testing.T) {
	is := is.New(t)

	input, err := os.Open("testdata/2017-meta.txt")
	is.NoErr(err)

	defer input.Close()

	result, err := nvd.DecodeMeta(input)
	is.NoErr(err)
	is.Equal(result.Size, 130731928)
	is.Equal(result.GzipSize, 7357077)
	is.Equal(result.Sha256, "6F961E24FD62E24165054CA78CBD5A8468420C893B15E118051E607BB4E3E7D9")
	is.True(result.LastModified.Equal(time.Date(2018, 9, 15, 7, 15, 33, 0, time.UTC)))
}
