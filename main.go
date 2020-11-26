//go:generate go run asset-generator.go
package main

import (
	"github.com/t-richards/vulnsearch/internal/cmd"
)

func main() {
	cmd.Execute()
}
