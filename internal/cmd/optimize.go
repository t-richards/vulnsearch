package cmd

import (
	"log"

	"github.com/t-richards/vulnsearch/internal/app"

	"github.com/spf13/cobra"
)

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize the database",
	Long:  "Optimize the database by running VACUUM",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Optimizing database...")
		app := app.New()
		app.DB.Exec("PRAGMA synchronous = OFF")
		app.DB.Exec("PRAGMA journal_mode = memory")
		app.DB.Exec("PRAGMA optimize")
		app.DB.Exec("VACUUM")
		log.Printf("Done.")
	},
}
