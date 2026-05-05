package stats

import "github.com/Judisk/daily-tracker-cli/internal/storage"

func AvgMood(records []storage.Record) float64 {
	if len(records) == 0 {
		return 0
	}

	sum := 0
	for _, r := range records {
		sum += r.Mood
	}

	return float64(sum) / float64(len(records))
}
