// Package assets embeds the front end source files, which the server compiles
// at startup.
package assets

import _ "embed"

// ApplicationTS is the TypeScript source for the search page.
//
//go:embed application.ts
var ApplicationTS string
