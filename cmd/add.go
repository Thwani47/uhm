package cmd

import (
	"fmt"
	"regexp"
	"strings"

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

		commands, err := database.ListCommands()

		if err != nil {
			fmt.Printf("Error getting commands: %v\n", err)
			return
		}

		availableCommands := make(map[string]promptutils.PromptSuggestion, len(commands))

		for _, command := range commands {
			availableCommands[command.Name] = promptutils.PromptSuggestion{
				Text:        command.Name,
				Description: command.Description,
				Value:       command.Command,
			}
		}

		commandName := promptutils.PromptInput(promptutils.PromptContent{
			Label:        "Enter the command name: ",
			ErrorMessage: "Command name cannot be empty",
		}, promptutils.AdditionalValidation{
			ErrorMesage: "the command name already exists",
			ValidationFunc: func(input string) bool {
				_, ok := availableCommands[input]

				return ok
			},
		})

		commandValue := promptutils.PromptInput(promptutils.PromptContent{
			Label:        "Enter the command: (Add variables with $variable_name)",
			ErrorMessage: "Command cannot be empty",
		}, promptutils.AdditionalValidation{ // we don't need any additional validtion for the command
			ErrorMesage: "",
			ValidationFunc: func(input string) bool {
				return false
			},
		})

		re := regexp.MustCompile(`\$(\w+)\b`)
		commandValue = re.ReplaceAllStringFunc(commandValue, func(s string) string {
			variableName := strings.TrimPrefix(s, "$")
			return "{{ ." + variableName + " }}"
		})

		commandDescription := promptutils.PromptInput(promptutils.PromptContent{
			Label:        "Enter a description for the command: ",
			ErrorMessage: "Description cannot be empty",
		}, promptutils.AdditionalValidation{ // we don't need any additional validtion for the description
			ErrorMesage: "",
			ValidationFunc: func(input string) bool {
				return false
			},
		})

		categories, err := database.ListCategories()
		categories = append(categories, "None")

		if err != nil {
			fmt.Printf("Error getting categories: %v\n", err)
			return
		}

		category := promptutils.PromptSelect(promptutils.SelectPrompt{
			Label:   "Select a category for the command: ",
			Options: categories,
		})

		if category == "None" {
			category = ""
		}

		err = database.AddCommand(commandName, commandValue, category, commandDescription)

		if err != nil {
			fmt.Printf("Error adding command: %v\n", err)
			return
		}

		fmt.Println("Command added successfully ✅")
	},
}
