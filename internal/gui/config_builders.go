package gui

func buildConfigs() []fieldConfig {
	return []fieldConfig{
		newScoreFieldConfigTime("Took meds"),
		newScoreFieldConfigTime("Went to bed"),
		newScoreFieldConfigTime("Fell asleep"),
		newScoreFieldConfigTime("Woke Up"),

		newScoreFieldConfigInt("Sleep Quality"),
		newScoreFieldConfigInt("Mood"),
		newScoreFieldConfigInt("Energy"),
		newScoreFieldConfigInt("Focus"),
		intFieldConfig{
			prompt: "Pills",
			Min:    0,
			Max:    50,
		},
	}
}

func newScoreFieldConfigInt(prompt string) intFieldConfig {
	return intFieldConfig{
		prompt: prompt,
		Min:    0,
		Max:    5,
	}
}

func newScoreFieldConfigTime(prompt string) timeFieldConfig {
	return timeFieldConfig{
		prompt: prompt,
	}
}
