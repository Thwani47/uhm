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

type PromptSelectRunner interface {
	Run() (int, string, error)
	SetupSelectPrompt(pc PromptContent, items []string)
}

type PromptInputRunner interface {
	Run() (string, error)
	SetupInputPrompt(pc PromptContent, validate func(string) error, templates *promptui.PromptTemplates)
}

type RealSelectRunner struct {
	Select promptui.Select
}

func (r *RealSelectRunner) Run() (int, string, error) {
	return r.Select.Run()
}

func (r *RealSelectRunner) SetupSelectPrompt(pc PromptContent, items []string) {
	r.Select = promptui.Select{
		Label: pc.Label,
		Items: items,
	}
}

type RealInputRunner struct {
	Prompt promptui.Prompt
}

func (r *RealInputRunner) Run() (string, error) {
	return r.Prompt.Run()
}

func (r *RealInputRunner) SetupInputPrompt(pc PromptContent, validate func(string) error, templates *promptui.PromptTemplates) {
	r.Prompt = promptui.Prompt{
		Label:     pc.Label,
		Validate:  validate,
		Templates: templates,
	}
}

type AdditionalValidation struct {
	ErrorMesage    string
	ValidationFunc func(string) bool
}

func PromptInput(pc PromptContent, addionalValidation AdditionalValidation, runner PromptInputRunner) string {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New(pc.ErrorMessage)
		}

		if addionalValidation.ValidationFunc(input) {
			return errors.New(addionalValidation.ErrorMesage)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}: ",
		Valid:   "{{ . | green }}: ",
		Invalid: "{{ . | red }}: ",
		Success: "{{ . | bold }}: ",
	}

	runner.SetupInputPrompt(pc, validate, templates)

	result, err := runner.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func PromptSelect(pc PromptContent, items []string, runner PromptSelectRunner) string {
	runner.SetupSelectPrompt(pc, items)

	_, result, err := runner.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
