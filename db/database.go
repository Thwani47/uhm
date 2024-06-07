package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func createOrOpenDatabase() (*DB, error) {
	dbPath := "uhm.db"

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
