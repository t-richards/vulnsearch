package cmd

import (
	"github.com/t-richards/vulnsearch/internal/nvd"

	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data to disk",
	Long:  "Fetch compressed CVE data archives from NVD",
	Run: func(cmd *cobra.Command, args []string) {
		nvd.DownloadAll()
	},
}
