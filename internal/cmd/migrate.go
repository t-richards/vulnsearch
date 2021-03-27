package cmd

import (
	"fmt"

	"github.com/t-richards/vulnsearch/internal/db"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Long:  "Migrate the database to the latest schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Migrating the database... ")
		conn := db.Connect()
		db.Migrate(conn)
		fmt.Print("Done.\n")
	},
}
