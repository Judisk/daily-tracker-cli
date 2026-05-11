package stats

import (
	"fmt"

	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/storage"
)

func Avg(records []model.Record, selector func(model.Record) int) float64 {
	if len(records) == 0 {
		return 0
	}

	sum := 0
	for _, r := range records {
		sum += selector(r)
	}

	return float64(sum) / float64(len(records))
}

func RunStats(last int) {
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

	avgSleepQuality := Avg(records, func(r model.Record) int { return r.SleepQuality })

	avgMood := Avg(records, func(r model.Record) int { return r.Mood })

	avgEnergy := Avg(records, func(r model.Record) int { return r.Energy })

	avgFocus := Avg(records, func(r model.Record) int { return r.Focus })

	fmt.Printf("Records used: %d\n", len(records))
	fmt.Printf("Average sleep quality:   %.2f\n", avgSleepQuality)
	fmt.Printf("Average mood:   %.2f\n", avgMood)
	fmt.Printf("Average energy: %.2f\n", avgEnergy)
	fmt.Printf("Average focus:  %.2f\n", avgFocus)
}
