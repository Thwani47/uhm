package cmd

import (
	"fmt"

	database "github.com/Thwani47/uhm/db"
	promptutils "github.com/Thwani47/uhm/internal/prompt_utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a command",
	Long:  "Delete a command from the list of commonly used commands",
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := database.ListCommands()

		if err != nil {
			fmt.Printf("Error getting commands: %v\n", err)
			return
		}

		availableCommands := make([]string, len(commands)+1)

		for index, command := range commands {
			availableCommands[index] = command.Name
		}

		availableCommands[len(commands)] = "Cancel"

		commandsToDelete := promptutils.PromptMultiSelect(promptutils.MultiSelectPrompt{
			Label:        "Select the commands to delete",
			ErrorMessage: "Please select at least one command",
			Options:      availableCommands,
		})

		if len(commandsToDelete) == 0 || contains(commandsToDelete, "Cancel") {
			fmt.Println("No commands selected for deletion")
			return
		}

		if !promptutils.PromptConfirm(promptutils.ConfirmPrompt{
			Title:       "Are you sure you want to delete the selected command(s)?",
			Affirmative: "Yes",
			Negative:    "No",
		}) {
			fmt.Println("Deletion cancelled")
			return
		}

		err = database.DeleteCommands(commandsToDelete)

		if err != nil {
			fmt.Printf("Error deleting command(s): %v\n", err)
			return
		}

		fmt.Println("Command(s) deleted successfully! ✅")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func contains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}

	return false
}
