package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func AddCommand(name, command, category, description string) error {
	db, err := createOrOpenDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO commands (name, command, category, description) VALUES (?, ?, ?, ?)", name, command, category, description)

	if err != nil {
		return errors.New("please run `uhm init` to initialize the database")
	}

	return nil
}

type Command struct {
	Name        string
	Command     string
	Category    string
	Description string
}

func ListCommands() ([]Command, error) {
	db, err := createOrOpenDatabase()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT name, command, category, description FROM commands")

	if err != nil {
		return nil, errors.New("please run `uhm init` to initialize the database")
	}

	defer rows.Close()

	var commands []Command

	for rows.Next() {
		var command Command
		var category sql.NullString
		var description sql.NullString
		err = rows.Scan(&command.Name, &command.Command, &category, &description)

		if err != nil {
			return nil, err
		}

		if category.Valid {
			command.Category = category.String
		} else {
			command.Category = ""
		}

		if description.Valid {
			command.Description = description.String
		} else {
			command.Description = ""
		}

		commands = append(commands, command)
	}

	return commands, nil

}

func DeleteCommands(commands []string) error {
	db, err := createOrOpenDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	commandStr := strings.Join(commands, "','")

	sqlStatement := fmt.Sprintf("DELETE FROM commands WHERE name IN ('%s')", commandStr)

	_, err = db.Exec(sqlStatement)

	if err != nil {
		return errors.New("please run `uhm init` to initialize the database")
	}

	return nil
}

func ListCategories() ([]string, error) {
	db, err := createOrOpenDatabase()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT name FROM categories")

	if err != nil {
		return nil, errors.New("please run `uhm init` to initialize the database")
	}

	defer rows.Close()

	var categories []string

	for rows.Next() {
		var category string
		err = rows.Scan(&category)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
