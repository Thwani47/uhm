package cmd

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	database "github.com/Thwani47/uhm/db"
	promptutils "github.com/Thwani47/uhm/internal/prompt_utils"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uhm",
	Short: "uhm helps you remember your commonly used commands",
	Long:  "A CLI tool that helps you store yor commonly used commands and run them with ease.",
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := database.ListCommands()

		if err != nil {
			fmt.Println("Error listing commands", err)
			os.Exit(1)
		}

		availableCommands := make(map[string]promptutils.PromptSuggestion, len(commands))

		for _, command := range commands {
			availableCommands[command.Name] = promptutils.PromptSuggestion{
				Text:        command.Name,
				Description: command.Description,
				Value:       command.Command,
			}
		}

		selectedCommand := promptutils.MakeSelection(availableCommands)

		re := regexp.MustCompile(`{{\s*\.\w+\s*}}`)
		matches := re.FindAllString(selectedCommand, -1)

		if len(matches) > 0 {
			template, err := template.New("command").Parse(selectedCommand)

			if err != nil {
				fmt.Println("Error parsing command", err)
				os.Exit(1)
			}

			variablesMap := make(map[string]string)
			for _, match := range matches {
				variableName := strings.TrimSpace(match[4 : len(match)-2])
				variableValue := promptutils.PromptInput(promptutils.PromptContent{
					Label:        fmt.Sprintf("Enter value for %s", variableName),
					ErrorMessage: "Value cannot be empty",
				}, promptutils.AdditionalValidation{
					ErrorMesage: "",
					ValidationFunc: func(input string) bool {
						return false
					},
				})

				variablesMap[variableName] = variableValue
			}

			var command bytes.Buffer
			err = template.Execute(&command, variablesMap)

			if err != nil {
				fmt.Println("Error executing template", err)
				os.Exit(1)
			}

			selectedCommand = command.String()
		}

		if err := clipboard.WriteAll(selectedCommand); err != nil {
			fmt.Println("Error copying command to clipboard", err)
			os.Exit(1)
		}

		fmt.Println("Command copied to clipboard 📋")

	},
}

func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
