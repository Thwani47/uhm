package promptutils

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMessage string
	Label        string
}

func PromptInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New(pc.ErrorMessage)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}: ",
		Valid:   "{{ . | green }}: ",
		Invalid: "{{ . | red }}: ",
		Success: "{{ . | bold }}: ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Validate:  validate,
		Templates: templates,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func PromptSelect(pc PromptContent, items []string) string {
	prompt := promptui.Select{
		Label: pc.Label,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
