package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func AddCommand(name, command, description string) error {
	db, err := createOrOpenDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO commands (name, command, description) VALUES (?, ?, ?)", name, command, description)

	if err != nil {
		return errors.New("please run `uhm init` to initialize the database")
	}

	return nil
}

type Command struct {
	Name        string
	Command     string
	Description string
}

func ListCommands() ([]Command, error) {
	db, err := createOrOpenDatabase()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT name, command, description FROM commands")

	if err != nil {
		return nil, errors.New("please run `uhm init` to initialize the database")
	}

	defer rows.Close()

	var commands []Command

	for rows.Next() {
		var command Command
		var description sql.NullString
		err = rows.Scan(&command.Name, &command.Command, &description)

		if err != nil {
			return nil, err
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
