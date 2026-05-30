package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func RunGui() {
	a := app.New()
	w := a.NewWindow("Health Tracker")

	fields := []*inputField{}
	configs := buildConfigs()

	for _, cfg := range configs {
		field := buildField(cfg)
		fields = append(fields, field)
	}

	button := buildSaveButton(w, fields)
	exportButton := buildExportButton(w)
	form := buildForm(fields)
	content := container.NewVBox(
		form,
		button,
		exportButton,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
