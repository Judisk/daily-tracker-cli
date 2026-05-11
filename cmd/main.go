package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Judisk/daily-tracker-cli/internal/input"
	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/sleep"
	"github.com/Judisk/daily-tracker-cli/internal/stats"
	"github.com/Judisk/daily-tracker-cli/internal/storage"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	add := flag.Bool("add", false, "Add new record")
	statsFlag := flag.Bool("stats", false, "Show stats")
	last := flag.Int("last", 0, "Show last N days")

	flag.Parse()

	if *add && *statsFlag {
		fmt.Println("Choose only one: --add or --stats")
		return
	}

	if !*add && !*statsFlag {
		fmt.Println("Usage:")
		fmt.Println("  --add    Add new record")
		fmt.Println("  --stats  Show stats")
		fmt.Println("  --stats --last 7")
		return
	}
	if *statsFlag {
		runStats(*last)
	} else {
		runAdd(reader)
	}
}

type Field[T any] struct {
	prompt   string
	Validate func(string) (T, error)
}

func runAdd(r *bufio.Reader) {

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

func runStats(last int) {
	records, err := storage.Load()
	if err != nil {
		fmt.Printf("failed to load data: %v\n", err)
		return
	}

	if len(records) == 0 {
		fmt.Println("No data yet")
		return
	}

	if last > 0 {
		if last < len(records) {
			records = records[len(records)-last:]
		}
	}

	avgSleepQuality := stats.Avg(records, func(r model.Record) int { return r.SleepQuality })

	avgMood := stats.Avg(records, func(r model.Record) int { return r.Mood })

	avgEnergy := stats.Avg(records, func(r model.Record) int { return r.Energy })

	avgFocus := stats.Avg(records, func(r model.Record) int { return r.Focus })

	fmt.Printf("Records used: %d\n", len(records))
	fmt.Printf("Average sleep quality:   %.2f\n", avgSleepQuality)
	fmt.Printf("Average mood:   %.2f\n", avgMood)
	fmt.Printf("Average energy: %.2f\n", avgEnergy)
	fmt.Printf("Average focus:  %.2f\n", avgFocus)
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
