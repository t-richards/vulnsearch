//go:generate go run asset-generator.go
package main

import (
	"log"
	"os"
	"vulnsearch/internal/app"
	"vulnsearch/internal/cli"
	"vulnsearch/internal/nvd"
)

var cmds cli.SubcommandSet

func help() {
	log.Printf("Valid subcommands: %v", cmds.ValidCommands())
}

func serve() {
	app := app.New()
	app.Run()
}

func main() {
	cmds = make(cli.SubcommandSet)
	cmds["fetch"] = nvd.DownloadAll
	cmds["help"] = help
	cmds["load"] = nvd.LoadFiles
	cmds["migrate"] = cli.Migrate
	cmds["optimize"] = cli.Optimize
	cmds["serve"] = serve

	if len(os.Args) < 2 {
		log.Fatalf("Please specify a subcommand: %v", cmds.ValidCommands())
	}

	err := cmds.Run(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}
}
