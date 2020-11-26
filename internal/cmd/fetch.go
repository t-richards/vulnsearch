package cmd

import (
	"vulnsearch/internal/nvd"

	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data to disk",
	Long:  "Fetches compressed CVE archives from the NVD site",
	Run: func(cmd *cobra.Command, args []string) {
		nvd.DownloadAll()
	},
}
