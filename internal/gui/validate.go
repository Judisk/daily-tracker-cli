package gui

import (
	"fmt"
	"strconv"
	"time"
)

func (f *inputField) validate() bool {
	return f.config.Validate(f)
}

func (c intFieldConfig) Validate(f *inputField) bool {
	value, err := strconv.Atoi(f.entry.Text)
	msg := fmt.Sprintf("%s must be %d-%d",
		c.prompt,
		c.Min,
		c.Max)
	if err != nil {
		f.errorLabel.SetText(msg)
		return false
	}
	if value < c.Min || value > c.Max {
		f.errorLabel.SetText(msg)
		return false
	}
	f.value = value
	f.errorLabel.SetText("")
	return true
}

func (c timeFieldConfig) Validate(f *inputField) bool {
	value, err := time.Parse("15:04", f.entry.Text)

	if err != nil {
		f.errorLabel.SetText("Incorrect time")
		return false
	}
	f.value = value
	f.errorLabel.SetText("")
	return true
}
