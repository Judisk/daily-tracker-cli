package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Judisk/daily-tracker-cli/internal/input"
	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/stats"
	"github.com/Judisk/daily-tracker-cli/internal/storage"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	add := flag.Bool("add", false, "Add new record")
	statsFlag := flag.Bool("stats", false, "Show stats")
	last := flag.Int("last", 0, "Show last N days")
	/* пока делаю без флагов потом придумаю как их внедрить в input
	mood := flag.Int("mood", -1, "Mood (0-5)")
	energy := flag.Int("energy", -1, "Energy (0-5)")
	focus := flag.Int("focus", -1, "Focus (0-5)")
	pills := flag.Int("pills", -1, "Pills (0-50)")
	*/
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
		runAddWithFlag()
	}
}

func runAddWithFlag() {

	mood, err := getValue(reader, "Mood 0-5", input.ParseAndValidateInt("mood", model.MinValue, model.MaxValue))
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	energy, err := getValue(reader, "Energy 0-5", input.ParseAndValidateInt("energy", model.MinValue, model.MaxValue))
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	focus, err := getValue(reader, "Focus 0-5", input.ParseAndValidateInt("focus", model.MinValue, model.MaxValue))
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	pills, err := getValue(reader, "Pills 0-50", input.ParseAndValidateInt("pills", model.MinValue, model.PillsMax))
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	if pills <= model.PillsLowThreshold {
		fmt.Printf("Warning pills running low (%d left)\n", pills)
	}

	record, err := storage.NewRecord(mood, energy, focus, pills)
	if err != nil {
		fmt.Println("Error:", err)
		return
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

	avgMood := stats.Avg(records, func(r storage.Record) int { return r.Mood })

	avgEnergy := stats.Avg(records, func(r storage.Record) int { return r.Energy })

	avgFocus := stats.Avg(records, func(r storage.Record) int { return r.Focus })

	fmt.Printf("Records used: %d\n", len(records))
	fmt.Printf("Average mood:   %.2f\n", avgMood)
	fmt.Printf("Average energy: %.2f\n", avgEnergy)
	fmt.Printf("Average focus:  %.2f\n", avgFocus)
}

func getValue[T any](r *bufio.Reader, prompt string, f func(string) (T, error)) (v T, err error) {

	fmt.Print(prompt)
	for {
		fmt.Print(" -> ")
		v, err := input.Input(r, f)
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

/*
func getValue(flagVal *int, prompt string, min, max int, f func(string)()) (int, error) {
	if *flagVal != -1 {
		if *flagVal < min || *flagVal > max {
			fmt.Println("Invalid value, must be between", min, max)
			return 0, fmt.Errorf("value must be between %d and %d", min, max)
		}
		return *flagVal, nil
	}
	for {
		v, err := input.Input(reader)
		if err != nil {

		}
	}
}*/
