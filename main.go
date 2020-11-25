package main

import (
	"fmt"
	"os"
	"vulnsearch/internal/nvd"
	"vulnsearch/internal/webserver"
)

func main() {
	// Configure subcommands
	subCommands := make(map[string]func())
	subCommands["fetch"] = nvd.DownloadAll
	subCommands["load"] = nvd.LoadFiles
	subCommands["serve"] = webserver.Serve

	if len(os.Args) < 2 {
		fmt.Print("Please specify a subcommand: ")
		for key := range subCommands {
			fmt.Printf("%v ", key)
		}
		fmt.Print("\n")

		os.Exit(1)
	}

	key := os.Args[1]
	cmd, ok := subCommands[key]
	if !ok {
		fmt.Printf("Invalid subcommand: %v", key)
		os.Exit(1)
	}

	cmd()
}
