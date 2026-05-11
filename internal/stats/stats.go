package stats

import "github.com/Judisk/daily-tracker-cli/internal/model"

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
