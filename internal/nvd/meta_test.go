package nvd_test

import (
	"os"
	"testing"

	"github.com/t-richards/vulnsearch/internal/nvd"

	"github.com/matryer/is"
)

func TestParseMeta(t *testing.T) {
	is := is.New(t)
	input, _ := os.Open("testdata/2017-meta.txt")

	result, err := nvd.DecodeMeta(input)

	is.Equal(err, nil)
	is.Equal(130731928, result.Size)
	is.Equal("6F961E24FD62E24165054CA78CBD5A8468420C893B15E118051E607BB4E3E7D9", result.Sha256)
}
