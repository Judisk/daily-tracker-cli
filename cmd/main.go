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
		runAdd()
	}
}

func runAdd() {
	mood := input.AskInt("Mood (0-5): ", 0, 5)
	energy := input.AskInt("Energy (0-5): ", 0, 5)
	focus := input.AskInt("Focus (0-5): ", 0, 5)

	record := storage.NewRecord(mood, energy, focus)
	storage.Save(record)

	fmt.Println("Saved ✅")
}

func runStats(last int) {
	records := storage.Load()

	if len(records) == 0 {
		fmt.Println("No data yet")
		return
	}

	if last > 0 && last < len(records) {
		records = records[len(records)-last:]
	}

	fmt.Println("Records used:", len(records))
	fmt.Println("Average mood:", stats.AvgMood(records))
}
