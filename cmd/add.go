package cmd

import (
	"fmt"

	database "github.com/Thwani47/uhm/db"
	promptutils "github.com/Thwani47/uhm/internal/prompt_utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new command",
	Long:  "Add a new command to the list of commonly used commands",
	Run: func(cmd *cobra.Command, args []string) {

		commandName := promptutils.PromptInput(promptutils.PromptContent{
			Label:        "Enter the command name: ",
			ErrorMessage: "Command name cannot be empty",
		}, &promptutils.RealInputRunner{})

		commandValue := promptutils.PromptInput(promptutils.PromptContent{
			Label:        "Enter the command: ",
			ErrorMessage: "Command cannot be empty",
		}, &promptutils.RealInputRunner{})

		commandDescription := promptutils.PromptInput(promptutils.PromptContent{
			Label:        "Enter a description for the command: ",
			ErrorMessage: "Description cannot be empty",
		}, &promptutils.RealInputRunner{})

		err := database.AddCommand(commandName, commandValue, commandDescription)

		if err != nil {
			fmt.Printf("Error adding command: %v\n", err)
			return
		}

		fmt.Println("Command added successfully ✅")
	},
}
