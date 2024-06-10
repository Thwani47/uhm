package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func createOrOpenDatabase() (*DB, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(homeDir, "uhm.db")

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	db, err := sql.Open("sqlite", dbPath)
	db.SetMaxOpenConns(1)

	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
