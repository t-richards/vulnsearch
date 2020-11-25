//go:generate go run asset-generator.go
package main

import (
	"vulnsearch/internal/cmd"
)

func main() {
	cmd.Execute()
}
