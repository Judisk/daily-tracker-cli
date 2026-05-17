package gui

import "fyne.io/fyne/v2/widget"

func buildForm(fields []*inputField) *widget.Form {
	var formItems []*widget.FormItem
	for _, field := range fields {
		formItems = append(formItems,
			widget.NewFormItem(field.config.Prompt(), field.content()),
		)
	}

	return widget.NewForm(formItems...)

}
