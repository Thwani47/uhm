package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new command",
	Long:  "Add a new command to the list of commonly used commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}
