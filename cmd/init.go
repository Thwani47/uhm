package cmd

import (
	database "github.com/Thwani47/uhm/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Long:  "Initialize the database for storing commands",
	Run: func(cmd *cobra.Command, args []string) {
		database.InitializeDatabase()
		database.CreateSchema()
	},
}
