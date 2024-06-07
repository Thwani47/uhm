package cmd

import (
	"fmt"
	"os"

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

		if err := clipboard.WriteAll(selectedCommand); err != nil {
			fmt.Println("Error copying command to clipboard", err)
			os.Exit(1)
		}

		fmt.Println("Command copied to clipboard 📋")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
