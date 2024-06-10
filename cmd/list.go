package cmd

import (
	"fmt"

	database "github.com/Thwani47/uhm/db"
	"github.com/fatih/color"
	"github.com/rodaine/table"
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

		if len(commands) == 0 {
			return
		}

		tbl := table.New("Name", "Description")
		tbl.WithHeaderFormatter(color.New(color.FgGreen, color.Bold).Sprintf).WithFirstColumnFormatter(color.New(color.FgYellow).Sprintf)

		for _, command := range commands {
			tbl.AddRow(command.Name, command.Description)
		}

		tbl.Print()
	},
}
