package promptutils

import (
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
)

type MockSelectRunner struct {
	Label string
	Items []string
}

func (m *MockSelectRunner) Run() (int, string, error) {
	return 0, "item1", nil
}

func (m *MockSelectRunner) SetupSelectPrompt(pc PromptContent, items []string) {
	m.Label = pc.Label
	m.Items = items
}

type MockInputRunner struct {
	Label        string
	ErrorMessage string
	ShouldError  bool
}

func (m *MockInputRunner) Run() (string, error) {
	if m.ShouldError {
		return "", nil
	}
	return "input", nil
}

func (m *MockInputRunner) SetupInputPrompt(pc PromptContent, validate func(string) error, templates *promptui.PromptTemplates) {
	m.Label = pc.Label
	m.ErrorMessage = pc.ErrorMessage
}

func TestPromptInput(t *testing.T) {
	pc := PromptContent{
		Label:        "enter an input",
		ErrorMessage: "Input cannot be empty",
	}

	result := PromptInput(pc, &MockInputRunner{
		ShouldError: false,
	})

	assert.Equal(t, "input", result)
}

func TestPromptInput_Error(t *testing.T) {
	pc := PromptContent{
		Label:        "enter an input",
		ErrorMessage: "Input cannot be empty",
	}

	result := PromptInput(pc, &MockInputRunner{
		ShouldError: true,
	})

	assert.Equal(t, "", result)
}

func TestPromptSelect(t *testing.T) {
	pc := PromptContent{
		Label:        "select an item",
		ErrorMessage: "Item cannot be empty",
	}

	items := []string{"item1", "item2"}

	result := PromptSelect(pc, items, &MockSelectRunner{})

	assert.Equal(t, "item1", result)
}
