package promptutils

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

type PromptContent struct {
	ErrorMessage string
	Label        string
}

type SelectPrompt struct {
	ErrorMessage string
	Label        string
	Options      []string
}

type AdditionalValidation struct {
	ErrorMesage    string
	ValidationFunc func(string) bool
}

type ConfirmPrompt struct {
	Title       string
	Affirmative string
	Negative    string
}

func PromptInput(pc PromptContent, additionalValidation AdditionalValidation) string {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New(pc.ErrorMessage)
		}

		if additionalValidation.ValidationFunc(input) {
			return errors.New(additionalValidation.ErrorMesage)
		}

		return nil
	}

	var result string

	err := huh.
		NewInput().
		Title(pc.Label).
		Validate(validate).
		Value(&result).Run()

	if err != nil {
		fmt.Print("Prompt failed ", err, "\n")
		os.Exit(1)
	}

	return result
}

func PromptMultiSelect(sp SelectPrompt) []string {
	var selections []string
	options := make([]huh.Option[string], len(sp.Options))

	for index, option := range sp.Options {
		options[index] = huh.Option[string]{
			Key:   fmt.Sprint(option),
			Value: option,
		}
	}
	huh.NewMultiSelect[string]().Title(sp.Label).Options(options...).Value(&selections).Run()

	return selections
}

func PromptSelect(sp SelectPrompt) string {
	var selection string
	options := make([]huh.Option[string], len(sp.Options))

	for index, option := range sp.Options {
		options[index] = huh.Option[string]{
			Key:   fmt.Sprint(option),
			Value: option,
		}
	}

	huh.NewSelect[string]().Title(sp.Label).Options(options...).Value(&selection).Run()

	return selection
}

func PromptConfirm(cp ConfirmPrompt) bool {
	var result bool

	err := huh.
		NewConfirm().
		Title(cp.Title).
		Affirmative(cp.Affirmative).
		Negative(cp.Negative).
		Value(&result).
		Run()

	if err != nil {
		fmt.Print("Prompt failed ", err, "\n")
		os.Exit(1)
	}

	return result
}
