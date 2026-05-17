package gui

type fieldConfig interface {
	Validate(*inputField) bool
	Prompt() string
}

type intFieldConfig struct {
	prompt string
	Min    int
	Max    int
}
type timeFieldConfig struct {
	prompt string
}

func (c intFieldConfig) Prompt() string {
	return c.prompt
}
func (c timeFieldConfig) Prompt() string {
	return c.prompt
}
