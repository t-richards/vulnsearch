package main

import (
	"fmt"
	"vulnsearch/internal/nvd"
)

func main() {
	result := nvd.DownloadMeta(2020)
	fmt.Printf("%+v", result)
}
