package gui

import "fyne.io/fyne/v2/widget"

func buildField(cfg fieldConfig) *inputField {
	field := inputField{
		entry:      widget.NewEntry(),
		errorLabel: widget.NewLabel(""),
		config:     cfg,
	}

	field.entry.OnChanged = func(string) {
		field.validate()
	}

	return &field
}
