package main

import (
	"flag"
	"fmt"

	"github.com/Judisk/daily-tracker-cli/internal/input"
	"github.com/Judisk/daily-tracker-cli/internal/stats"
	"github.com/Judisk/daily-tracker-cli/internal/storage"
)

func main() {

	add := flag.Bool("add", false, "Add new record")
	statsFlag := flag.Bool("stats", false, "Show stats")
	last := flag.Int("last", 0, "Show last N days")
	mood := flag.Int("mood", -1, "Mood (0-5)")
	energy := flag.Int("energy", -1, "Energy (0-5)")
	focus := flag.Int("focus", -1, "Focus (0-5)")
	pills := flag.Int("pills", -1, "Pills (0-50)")

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
		runAddWithFlag(mood, energy, focus, pills)
	}
}

func runAddWithFlag(moodF, energyF, focusF, pillsF *int) {

	mood, err := getValue(moodF, "Mood (0-5): ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	energy, err := getValue(energyF, "Energy (0-5): ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	focus, err := getValue(focusF, "Focus (0-5): ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	pills, err := getValuePills(pillsF, "Pills (0-50):")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	record, err := storage.NewRecord(mood, energy, focus, pills)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	storage.Save(record)
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

func getValue(flagVal *int, prompt string) (int, error) {
	if *flagVal != -1 {
		if *flagVal < storage.MinValue || *flagVal > storage.MaxValue {
			fmt.Println("Invalid value, must be between", storage.MinValue, storage.MaxValue)
			return 0, fmt.Errorf("value must be between %d and %d", storage.MinValue, storage.MaxValue)
		}
		return *flagVal, nil
	}
	return input.AskInt(prompt, storage.MinValue, storage.MaxValue)
}

func getValuePills(flagVal *int, prompt string) (int, error) {
	if *flagVal != -1 {
		if *flagVal < storage.PillsMin || *flagVal > storage.PillsMax {
			fmt.Println("Invalid value, must be between", storage.PillsMin, storage.PillsMax)
			return 0, fmt.Errorf("value must be between %d and %d", storage.PillsMin, storage.PillsMax)
		}
		return *flagVal, nil
	}
	input, err := input.AskInt(prompt, storage.PillsMin, storage.PillsMax)
	if input <= storage.PillsLowThreshold {
		fmt.Printf("Warning pills running low (%d left)\n", input)
	}
	return input, err
}
