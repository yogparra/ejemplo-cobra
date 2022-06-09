package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ejemplo-cobra",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.00.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	//versionCmd.Flags().BoolP("version", "v", false, "version number of ejemplo-cobra")
}
