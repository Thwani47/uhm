package cmd

import (
	"database/sql"
	"fmt"
	"os"

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
		db := createOrOpenDatabase()
		createSchema(db)
		defer db.Close()
	},
}

func createSchema(db *sql.DB) {
	fmt.Println("Creating schema...⌛")

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS commands (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		command TEXT NOT NULL
	)`)

	if err != nil {
		fmt.Printf("Error creating schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Schema created ✅")

}

func createOrOpenDatabase() *sql.DB {
	fmt.Println("Initializing database...⌛")
	db, err := database.OpenDb()

	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Database initialized ✅")
	return db
}
