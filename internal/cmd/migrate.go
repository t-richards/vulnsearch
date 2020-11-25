package cmd

import (
	"fmt"
	"vulnsearch/internal/app"
	"vulnsearch/internal/db"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrates the database",
	Long:  "Migrates the database to the latest schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Migrating the database... ")
		app := app.New()
		app.DB.Exec(db.Schema)
		fmt.Print("Done.\n")
	},
}
