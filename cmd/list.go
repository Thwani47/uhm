package cmd

import (
	"fmt"

	database "github.com/Thwani47/uhm/db"
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
		db, err := database.OpenDb()

		if err != nil {
			fmt.Printf("Error opening database: %v\n", err)
			return
		}

		defer db.Close()

		rows, err := db.Query("SELECT * FROM commands")

		if err != nil {
			fmt.Printf("Error querying database: %v\n", err)
			return
		}

		defer rows.Close()

		fmt.Println("Commands:")
		for rows.Next() {
			var id int
			var name, command string
			err = rows.Scan(&id, &name, &command)
			if err != nil {
				fmt.Printf("Error scanning row: %v\n", err)
				return
			}
			fmt.Printf("%d: %s: %s\n", id, name, command)
		}

	},
}
