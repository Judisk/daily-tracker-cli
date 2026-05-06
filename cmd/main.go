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
		runAddWithFlag(mood, energy, focus)
	}
}

/*
	func runAdd() {
		mood := input.AskInt("Mood (0-5): ", 0, 5)
		energy := input.AskInt("Energy (0-5): ", 0, 5)
		focus := input.AskInt("Focus (0-5): ", 0, 5)

		record := storage.NewRecord(mood, energy, focus)
		storage.Save(record)

}
*/
func runAddWithFlag(moodF, energyF, focusF *int) {
	var mood, energy, focus int

	if *moodF != -1 {
		mood = *moodF
	} else {
		mood = input.AskInt("Mood (0-5): ", 0, 5)
	}
	if *energyF != -1 {
		energy = *energyF
	} else {
		energy = input.AskInt("Energy (0-5): ", 0, 5)
	}

	if *focusF != -1 {
		focus = *focusF
	} else {
		focus = input.AskInt("Focus (0-5): ", 0, 5)
	}
	record, err := storage.NewRecord(mood, energy, focus)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	storage.Save(record)
	fmt.Println("Saved ✅")
}

func runStats(last int) {
	records := storage.Load()

	if len(records) == 0 {
		fmt.Println("No data yet")
		return
	}

	if last > 0 {
		if last < len(records) {
			records = records[len(records)-last:]
		}
	}

	fmt.Println("Records used:", len(records))
	fmt.Println("Average mood:", stats.AvgMood(records))
}
