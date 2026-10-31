// Package vulnsearch is a fast, offline-capable CVE search tool backed by the
// NVD data feeds.
package main

import "github.com/t-richards/vulnsearch/internal/cmd"

// Primary entry point for the application.
func main() {
	cmd.Execute()
}
