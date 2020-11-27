package cmd

import (
	"github.com/spf13/cobra"
	"github.com/t-richards/vulnsearch/internal/app"
	"github.com/t-richards/vulnsearch/internal/db"
	"github.com/t-richards/vulnsearch/internal/nvd"
)

var dbgCmd = &cobra.Command{
	Use: "dbg",
	Run: func(cmd *cobra.Command, args []string) {
		// setup()
		// test()
	},
}

func setup() {
	app := app.New()
	db.Migrate(app.DB)
	db.FastJournal(app.DB)
	nvd.LoadProducts(app.DB)
	db.Optimize(app.DB)
}

func test() {
	app := app.New()
	db.FastJournal(app.DB)
	nvd.LoadCVEsForYear(app.DB, 2020)
}
