package nvd_test

import (
	"os"
	"testing"

	"github.com/t-richards/vulnsearch/internal/nvd"

	"github.com/stretchr/testify/assert"
)

func TestParseMeta(t *testing.T) {
	input, _ := os.Open("testdata/2017-meta.txt")

	result, err := nvd.DecodeMeta(input)

	assert.Nil(t, err, "Meta parser returns without error")
	assert.Equal(t, 130731928, result.Size, "Meta parser parses the size")
	assert.Equal(t, "6F961E24FD62E24165054CA78CBD5A8468420C893B15E118051E607BB4E3E7D9", result.Sha256, "Meta parser parses the hash")
}
