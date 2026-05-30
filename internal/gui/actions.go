package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Judisk/daily-tracker-cli/internal/export"
)

func buildSaveButton(w fyne.Window, fields []*inputField) *widget.Button {
	return widget.NewButton("Save", func() {
		valid := true
		for _, field := range fields {
			if !field.validate() {
				valid = false
			}
		}
		if !valid {
			dialog.ShowInformation("Validation error", "Please fix the highlighted fields", w)
			return
		}

		if err := saveGui(fields); err != nil {
			dialog.ShowError(err, w)
			return
		}

		dialog.ShowInformation(
			"Saved",
			"Record saved successfully.",
			w,
		)
		clearFields(fields)

	})

}

func buildExportButton(w fyne.Window) *widget.Button {
	return widget.NewButton("Export", func() {
		if err := export.ExportJsonToCsv(); err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowInformation(
			"Exported",
			"CSV file exported successfully.",
			w,
		)
	})
}
