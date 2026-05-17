package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var errorSize = fyne.NewSize(200, 20)

type inputField struct {
	entry      *widget.Entry
	errorLabel *widget.Label
	config     fieldConfig
	value      any
}

func (f *inputField) content() *fyne.Container {
	entryContainer := container.NewGridWrap(
		fyne.NewSize(90, 35),
		f.entry,
	)

	errorContainer := container.NewGridWrap(
		errorSize,
		f.errorLabel,
	)
	return container.NewHBox(
		entryContainer,
		errorContainer,
	)

}
