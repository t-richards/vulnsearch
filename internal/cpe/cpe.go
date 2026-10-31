// Package cpe parses CPE 2.3 names and decides whether a product version falls
// within the version constraints of an NVD CPE match criterion.
package cpe

import (
	"errors"
	"fmt"
	"strings"
)

// Special attribute values from the CPE 2.3 naming specification.
const (
	Any           = "*" // Any value.
	NotApplicable = "-" // No meaningful value, e.g. hardware without versions.
)

// formattedPrefix begins every CPE 2.3 formatted string binding.
const formattedPrefix = "cpe:2.3:"

// componentCount is the number of attributes after the "cpe:2.3:" prefix.
const componentCount = 11

// ErrInvalidName is returned for strings which aren't CPE 2.3 formatted strings.
var ErrInvalidName = errors.New("invalid CPE 2.3 name")

// Name holds the attributes of a CPE 2.3 name used for product lookups.
type Name struct {
	Part    string // "a" (application), "o" (operating system) or "h" (hardware).
	Vendor  string
	Product string
	Version string
	Update  string
}

// Parse parses a CPE 2.3 formatted string such as
// "cpe:2.3:a:apache:http_server:2.4.49:*:*:*:*:*:*:*".
//
// Quoted characters (e.g. "notepad\+\+") are unescaped. Components are split
// on unescaped colons only, so "\:" may appear within a component.
func Parse(s string) (Name, error) {
	rest, ok := strings.CutPrefix(s, formattedPrefix)
	if !ok {
		return Name{}, fmt.Errorf("%w: %q", ErrInvalidName, s)
	}

	parts := split(rest)
	if len(parts) != componentCount {
		return Name{}, fmt.Errorf("%w: %q has %d components", ErrInvalidName, s, len(parts))
	}

	return Name{
		Part:    parts[0],
		Vendor:  parts[1],
		Product: parts[2],
		Version: parts[3],
		Update:  parts[4],
	}, nil
}

// split splits s on unescaped colons and unescapes each component.
func split(s string) []string {
	var (
		parts   []string
		current strings.Builder
		escaped bool
	)

	for _, r := range s {
		switch {
		case escaped:
			current.WriteRune(r)

			escaped = false
		case r == '\\':
			escaped = true
		case r == ':':
			parts = append(parts, current.String())
			current.Reset()
		default:
			current.WriteRune(r)
		}
	}

	return append(parts, current.String())
}
