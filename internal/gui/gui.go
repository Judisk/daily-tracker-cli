package gui

import (
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	application "github.com/Judisk/daily-tracker-cli/internal/app"
)

func RunGui(service *application.Service) {
	a := fyneapp.New()
	w := a.NewWindow("Health Tracker")

	fields := []*inputField{}
	configs := buildConfigs()

	for _, cfg := range configs {
		field := buildField(cfg)
		fields = append(fields, field)
	}

	button := buildSaveButton(w, fields, service)
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
