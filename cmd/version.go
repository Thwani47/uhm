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
	Short: "Print uhm version",
	Long:  "Print uhm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uhm v0.0.1")
	},
}
