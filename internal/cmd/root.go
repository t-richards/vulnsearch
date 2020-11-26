package cmd

import (
	"fmt"
	"os"

	"github.com/t-richards/vulnsearch/internal/app"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(optimizeCmd)
	rootCmd.AddCommand(fetchCmd)
}

var rootCmd = &cobra.Command{
	Use:   "github.com/t-richards/vulnsearch",
	Short: "github.com/t-richards/vulnsearch is a fast vulnerability search tool",
	Long:  "A fast, well-behaved replacement for other CVE search tools.",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.New()
		app.Run()
	},
}

// Execute runs the main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
