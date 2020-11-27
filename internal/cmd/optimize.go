package cmd

import (
	"log"

	"github.com/t-richards/vulnsearch/internal/app"
	"github.com/t-richards/vulnsearch/internal/db"

	"github.com/spf13/cobra"
)

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize the database",
	Long:  "Optimize the database by running VACUUM",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Optimizing database...")
		app := app.New()
		db.Optimize(app.DB)
		log.Printf("Done.")
	},
}
