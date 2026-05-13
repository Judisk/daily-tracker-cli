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

func getValue[T any](r *bufio.Reader, str formField[T]) (v T, err error) {

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

func newRecord(r *bufio.Reader, f fields) (model.Record, error) {

	wentToBed, err := getValue(r, f.wentToBed)
	if err != nil {

		return model.Record{}, fmt.Errorf("get went to bed: %w", err)
	}

	fellAsleep, err := getValue(r, f.fellAsleep)
	if err != nil {
		return model.Record{}, fmt.Errorf("get fell asleep: %w", err)
	}
	wokeUp, err := getValue(r, f.wokeUp)
	if err != nil {
		return model.Record{}, fmt.Errorf("get woke up: %w", err)
	}
	sleepQuality, err := getValue(r, f.sleepQuality)
	if err != nil {
		return model.Record{}, fmt.Errorf("get sleep quality: %w", err)
	}
	mood, err := getValue(r, f.mood)
	if err != nil {
		return model.Record{}, fmt.Errorf("get mood :%w", err)
	}
	energy, err := getValue(r, f.energy)
	if err != nil {
		return model.Record{}, fmt.Errorf("get energy: %w", err)
	}
	focus, err := getValue(r, f.focus)
	if err != nil {
		return model.Record{}, fmt.Errorf("get focus: %w", err)
	}
	tookMeds, err := getValue(r, f.tookMeds)
	if err != nil {
		return model.Record{}, fmt.Errorf("get took meds: %w", err)
	}
	pills, err := getValue(r, f.pills)
	if err != nil {
		return model.Record{}, fmt.Errorf("get pills: %w", err)
	}
	notes, err := getValue(r, f.notes)
	if err != nil {
		return model.Record{}, fmt.Errorf("get notes: %w", err)
	}
	sideEffects, err := getValue(r, f.sideEffects)
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
