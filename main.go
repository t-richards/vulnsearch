package main

import (
	"log"
	"os"
	"vulnsearch/internal/cli"
	"vulnsearch/internal/db"
	"vulnsearch/internal/nvd"
	"vulnsearch/internal/webserver"
)

func main() {
	cmds := make(cli.SubcommandSet)
	cmds["fetch"] = nvd.DownloadAll
	cmds["load"] = nvd.LoadFiles
	cmds["migrate"] = db.Migrate
	cmds["optimize"] = db.Optimize
	cmds["serve"] = webserver.Serve

	if len(os.Args) < 2 {
		log.Fatalf("Please specify a subcommand: %v", cmds.ValidCommands())
	}

	err := cmds.Run(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}
}
