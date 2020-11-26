package cmd

import (
	"github.com/t-richards/vulnsearch/internal/nvd"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load archive data",
	Long:  "Load compressed, already-downloaded archives from the cache directory into the database",
	Run: func(cmd *cobra.Command, args []string) {
		nvd.LoadFiles()
	},
}
