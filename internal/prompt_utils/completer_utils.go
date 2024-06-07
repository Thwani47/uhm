package promptutils

import "github.com/c-bata/go-prompt"

type PromptSuggestion struct {
	Text        string
	Description string
	Value       string
}

func MakeSelection(options map[string]PromptSuggestion) string {
	completer := func(d prompt.Document) []prompt.Suggest {
		var suggestions = []prompt.Suggest{}

		for _, option := range options {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        option.Text,
				Description: option.Description,
			})
		}

		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
	selection := prompt.Input("> ", completer)

	value, ok := options[selection]

	if !ok {
		return ""
	}

	return value.Value
}
