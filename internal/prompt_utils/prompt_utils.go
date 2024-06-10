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
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
