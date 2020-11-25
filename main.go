package main

import (
	"fmt"
	"os"
	"vulnsearch/internal/db"
	"vulnsearch/internal/nvd"
	"vulnsearch/internal/webserver"
)

type subCommands map[string]func()

func main() {
	cmds := make(subCommands)
	cmds["fetch"] = nvd.DownloadAll
	cmds["load"] = nvd.LoadFiles
	cmds["migrate"] = db.Migrate
	cmds["optimize"] = db.Optimize
	cmds["serve"] = webserver.Serve

	validateArgs(cmds)
	runSubcommand(cmds)
}

func validateArgs(cmds subCommands) {
	if len(os.Args) < 2 {
		fmt.Print("Please specify a subcommand: ")
		for key := range cmds {
			fmt.Printf("%v ", key)
		}
		fmt.Print("\n")

		os.Exit(1)
	}
}

func runSubcommand(cmds subCommands) {
	key := os.Args[1]
	cmd, ok := cmds[key]
	if !ok {
		fmt.Printf("Invalid subcommand: %v", key)
		os.Exit(1)
	}

	cmd()
}
