package gui

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
	"github.com/Judisk/daily-tracker-cli/internal/export"
)

func buildSaveButton(fields []*inputField) *widget.Button {
	return widget.NewButton("Save", func() {
		valid := true
		for _, field := range fields {
			if !field.validate() {
				valid = false
			}
		}
		if !valid {
			return
		}
		if err := saveGui(fields); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("save")
		}

	})

}

func buildExportButton() *widget.Button {
	return widget.NewButton("Export", func() {
		if err := export.ExportJsonToCsv(); err != nil {
			fmt.Println(err)
			return
		}
	})
}
