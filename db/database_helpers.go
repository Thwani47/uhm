package database

import (
	"fmt"
	"os"
)

func InitializeDatabase() {
	fmt.Println("Initializing database...⌛")
	db, err := createOrOpenDatabase()

	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()
	fmt.Println("Database initialized ✅")
}

func CreateSchema() {
	db, err := createOrOpenDatabase()

	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	fmt.Println("Creating schema...⌛")
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS commands (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		command TEXT NOT NULL,
		description TEXT NULL
	)`)

	if err != nil {
		fmt.Printf("Error creating database schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Schema created ✅")
}
