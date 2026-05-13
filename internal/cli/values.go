package cli

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/Judisk/daily-tracker-cli/internal/input"
	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/sleep"
)

func getValue[T any](r *bufio.Reader, str Field[T]) (v T, err error) {

	fmt.Print(str.prompt)
	for {
		fmt.Print(" -> ")
		v, err := input.Input(r, str.Validate)
		if err != nil {
			if err == io.EOF {
				return v, err
			}
			fmt.Println(err)
			continue
		}
		return v, nil
	}
}

func NewRecord(r *bufio.Reader, f Fields) (model.Record, error) {

	wentToBed, err := getValue(r, f.WentToBed)
	if err != nil {

		return model.Record{}, fmt.Errorf("get went to bed: %w", err)
	}

	fellAsleep, err := getValue(r, f.FellAsleep)
	if err != nil {
		return model.Record{}, fmt.Errorf("get fell asleep: %w", err)
	}
	wokeUp, err := getValue(r, f.WokeUp)
	if err != nil {
		return model.Record{}, fmt.Errorf("get woke up: %w", err)
	}
	sleepQuality, err := getValue(r, f.SleepQuality)
	if err != nil {
		return model.Record{}, fmt.Errorf("get sleep quality: %w", err)
	}
	mood, err := getValue(r, f.Mood)
	if err != nil {
		return model.Record{}, fmt.Errorf("get mood %w", err)
	}
	energy, err := getValue(r, f.Energy)
	if err != nil {
		return model.Record{}, fmt.Errorf("get energy: %w", err)
	}
	focus, err := getValue(r, f.Focus)
	if err != nil {
		return model.Record{}, fmt.Errorf("get focus: %w", err)
	}
	tookMeds, err := getValue(r, f.TookMeds)
	if err != nil {
		return model.Record{}, fmt.Errorf("get took meds: %w", err)
	}
	pills, err := getValue(r, f.Pills)
	if err != nil {
		return model.Record{}, fmt.Errorf("get pills %w", err)
	}
	notes, err := getValue(r, f.Notes)
	if err != nil {
		return model.Record{}, fmt.Errorf("get motes: %w", err)
	}
	sideEffects, err := getValue(r, f.SideEffects)
	if err != nil {
		return model.Record{}, fmt.Errorf("get side effects: %w", err)
	}

	sleepDuration := sleep.Duration(wokeUp, fellAsleep)

	return model.Record{
		Date:          time.Now(),
		WentToBed:     wentToBed,
		FellAsleep:    fellAsleep,
		WokeUp:        wokeUp,
		SleepDuration: sleepDuration,

		SleepQuality: sleepQuality,
		Mood:         mood,
		Energy:       energy,
		Focus:        focus,

		Pills:    pills,
		TookMeds: tookMeds,

		Notes:       notes,
		SideEffects: sideEffects,
	}, nil
}
