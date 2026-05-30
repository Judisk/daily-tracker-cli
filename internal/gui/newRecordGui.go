package gui

import (
	"fmt"
	"time"

	application "github.com/Judisk/daily-tracker-cli/internal/app"
	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/sleep"
)

func saveGui(fields []*inputField, service *application.Service) error {
	if err := service.CreateRecord(collectRecord(fields)); err != nil {
		return fmt.Errorf("create record: %w", err)
	}

	return nil
}

func collectRecord(fields []*inputField) model.Record {
	values := map[string]any{}

	for _, f := range fields {
		values[f.config.Prompt()] = f.value
	}

	sleepDuration := sleep.Duration(
		values["Woke Up"].(time.Time),
		values["Fell asleep"].(time.Time),
	)
	return model.Record{
		Date:          time.Now(),
		WentToBed:     values["Went to bed"].(time.Time),
		FellAsleep:    values["Fell asleep"].(time.Time),
		WokeUp:        values["Woke Up"].(time.Time),
		TookMeds:      values["Took meds"].(time.Time),
		SleepDuration: sleepDuration,

		SleepQuality: values["Sleep Quality"].(int),
		Mood:         values["Mood"].(int),
		Energy:       values["Energy"].(int),
		Focus:        values["Focus"].(int),
		Pills:        values["Pills"].(int),
	}
}
