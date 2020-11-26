package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/t-richards/vulnsearch/internal/cache"
)

const clobberMsg = `
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    DANGER: IRREVOCABLE DATA LOSS!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

`

var clobberCmd = &cobra.Command{
	Use:   "clobber",
	Short: "Clobber database and cache",
	Long:  "Completely removes the application's cache directory, which contains the sqlite database and all cache files. WARNING! This action is not reversible.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(clobberMsg)
		fmt.Print("This action will delete the following directory and its contents:\n\n")
		fmt.Printf("    %v\n", cache.CacheDir)
		fmt.Print("\nDo you wish to continue? [y/N]: ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(text)) == "y" {
			cache.Clobber()
			return
		}

		fmt.Println("Cancelled.")
	},
}
