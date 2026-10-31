package cpe_test

import (
	"errors"
	"testing"

	"github.com/matryer/is"

	"github.com/t-richards/vulnsearch/internal/cpe"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in   string
		want cpe.Name
	}{
		{
			"cpe:2.3:a:apache:http_server:2.4.49:*:*:*:*:*:*:*",
			cpe.Name{Part: "a", Vendor: "apache", Product: "http_server", Version: "2.4.49", Update: "*"},
		},
		{
			`cpe:2.3:a:notepad-plus-plus:notepad\+\+:8.7.5:*:*:*:*:*:*:*`,
			cpe.Name{Part: "a", Vendor: "notepad-plus-plus", Product: "notepad++", Version: "8.7.5", Update: "*"},
		},
		{
			`cpe:2.3:a:ruby-lang:cgi:*:*:*:*:*:ruby:*:*`,
			cpe.Name{Part: "a", Vendor: "ruby-lang", Product: "cgi", Version: "*", Update: "*"},
		},
		{
			`cpe:2.3:o:example:fw\:x:1\.0:sp1:*:*:*:*:*:*`,
			cpe.Name{Part: "o", Vendor: "example", Product: "fw:x", Version: "1.0", Update: "sp1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			is := is.New(t)

			got, err := cpe.Parse(tt.in)
			is.NoErr(err)
			is.Equal(got, tt.want)
		})
	}
}

func TestParseInvalid(t *testing.T) {
	for _, in := range []string{
		"",
		"cpe:/a:apache:http_server:2.4.49",
		"cpe:2.3:a:apache:http_server",
	} {
		t.Run(in, func(t *testing.T) {
			is := is.New(t)

			_, err := cpe.Parse(in)
			is.True(errors.Is(err, cpe.ErrInvalidName))
		})
	}
}
