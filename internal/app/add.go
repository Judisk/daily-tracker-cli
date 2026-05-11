package app

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/Judisk/daily-tracker-cli/internal/input"
	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/sleep"
	"github.com/Judisk/daily-tracker-cli/internal/storage"
)

type Field[T any] struct {
	prompt   string
	Validate func(string) (T, error)
}

func RunAdd(r *bufio.Reader) {

	wentToBedField := newTimeField("Went to bed")
	fellAsleepField := newTimeField("Fell asleep")
	wokeUpField := newTimeField("Woke up")
	sleepQualityField := newIntField("Sleep quality", model.MinValue, model.MaxValue)

	moodField := newIntField("Mood", model.MinValue, model.MaxValue)
	energyField := newIntField("Energy", model.MinValue, model.MaxValue)
	focusField := newIntField("Focus", model.MinValue, model.MaxValue)

	tookMedsField := newTimeField("Took meds")
	pillsField := newIntField("Pills", model.MinValue, model.PillsMax)

	notesField := newStringField("Notes")
	sideEffectsField := newStringField("Side Effects")

	wentToBed, err := getValue(r, wentToBedField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fellAsleep, err := getValue(r, fellAsleepField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	wokeUp, err := getValue(r, wokeUpField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	sleepQuality, err := getValue(r, sleepQualityField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	mood, err := getValue(r, moodField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	energy, err := getValue(r, energyField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	focus, err := getValue(r, focusField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	tookMeds, err := getValue(r, tookMedsField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	pills, err := getValue(r, pillsField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	notes, err := getValue(r, notesField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	sideEffects, err := getValue(r, sideEffectsField)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	sleepDuration := sleep.Sleep(wokeUp, fellAsleep)

	record := model.Record{
		Date: time.Now(),

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
	}

	if pills <= model.PillsLowThreshold {
		fmt.Printf("Warning pills running low (%d left)\n", pills)
	}

	if err := storage.Save(record); err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	fmt.Println("Saved ✅")
}

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

func newIntField(prompt string, min, max int) Field[int] {
	field := Field[int]{
		prompt:   fmt.Sprintf("%s %d-%d", prompt, min, max),
		Validate: input.IntValidator(prompt, min, max),
	}

	return field
}
func newTimeField(prompt string) Field[time.Time] {
	field := Field[time.Time]{
		prompt:   prompt,
		Validate: input.TimeValidator(),
	}

	return field
}
func newStringField(prompt string) Field[string] {
	field := Field[string]{
		prompt:   prompt,
		Validate: input.StringValidation(),
	}

	return field
}
