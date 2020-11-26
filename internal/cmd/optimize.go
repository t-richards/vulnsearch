package cmd

import (
	"log"

	"github.com/t-richards/vulnsearch/internal/app"

	"github.com/spf13/cobra"
)

const optimizeQuery = `
PRAGMA synchronous = OFF;
PRAGMA journal_mode = memory;
PRAGMA optimize;
VACUUM;
`

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize the database",
	Long:  "Optimize the database by running VACUUM",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Optimizing database...")
		app := app.New()
		app.DB.Exec(optimizeQuery)
		log.Printf("Done.")
	},
}
