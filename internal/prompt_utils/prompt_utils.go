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

type MultiSelectPrompt struct {
	ErrorMessage string
	Label        string
	Options      []string
}

type AdditionalValidation struct {
	ErrorMesage    string
	ValidationFunc func(string) bool
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

func PromptMultiSelect(pms MultiSelectPrompt) []string {
	var selections []string
	options := make([]huh.Option[string], len(pms.Options))

	for index, option := range pms.Options {
		options[index] = huh.Option[string]{
			Key:   fmt.Sprint(option),
			Value: option,
		}
	}
	huh.NewMultiSelect[string]().Title(pms.Label).Options(options...).Value(&selections).Run()

	return selections
}
