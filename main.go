package main

import (
	"log"
	"os"
	"vulnsearch/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Please specify a subcommand")
	}

	subCommand := os.Args[1]
	switch subCommand {
	case "fetch":
		commands.Fetch()
	case "load":
		commands.Load()
	default:
		log.Fatalf("Invalid subcommand: %v", subCommand)
	}
}
