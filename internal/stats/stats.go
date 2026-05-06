package stats

import "github.com/Judisk/daily-tracker-cli/internal/storage"

func Avg(records []storage.Record, get func(storage.Record) int) float64 {
	if len(records) == 0 {
		return 0
	}

	sum := 0
	for _, r := range records {
		sum += get(r)
	}

	return float64(sum) / float64(len(records))
}
