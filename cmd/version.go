package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of uhm",
	Long:  "Print the version number of uhm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uhm v0.0.1")
	},
}
