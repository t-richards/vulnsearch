package cmd

import (
	"fmt"
	"os"

	"github.com/t-richards/vulnsearch/internal/app"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "vulnsearch",
	Long: "vulnsearch: A fast, well-behaved replacement for other CVE search tools.",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.New()
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(optimizeCmd)
}

// Execute runs the main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
