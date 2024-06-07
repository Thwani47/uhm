package cmd

import (
	"fmt"

	database "github.com/Thwani47/uhm/db"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all commands",
	Long:  "List all stored commands",
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := database.ListCommands()

		if err != nil {
			fmt.Printf("Error listing commands: %v\n", err)
			return
		}

		for _, command := range commands {
			d := color.New(color.FgGreen, color.Bold)
			d.Printf("%s: ", command.Name)
			d = color.New(color.FgWhite).Add(color.Underline)
			d.Printf("%s", command.Command)
			fmt.Printf("  %s\n", command.Description)
		}
	},
}
