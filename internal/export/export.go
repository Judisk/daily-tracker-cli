package export

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/Judisk/daily-tracker-cli/internal/storage"
)

func ExportJsonToCsv() error {

	records, err := storage.Load()
	if err != nil {
		return err
	}

	file, err := os.Create("data/data.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	if err := writer.Write([]string{"Date", "Went to bed", "Fell asleep", "Woke up", "Sleep duration",
		"Sleep quality", "Mood", "Energy", "Focus",
		"Pills", "Took meds",
		"Notes", "Side effects",
	}); err != nil {
		return err
	}
	for _, r := range records {
		if err := writer.Write([]string{
			r.Date.Format("2006-01-02"),
			r.WentToBed.Format("15:04"),
			r.FellAsleep.Format("15:04"),
			r.WokeUp.Format("15:04"),
			r.SleepDuration.String(),
			strconv.Itoa(r.SleepQuality),
			strconv.Itoa(r.Mood),
			strconv.Itoa(r.Energy),
			strconv.Itoa(r.Focus),

			strconv.Itoa(r.Pills),
			r.TookMeds.Format("15:04"),
			r.Notes,
			r.SideEffects}); err != nil {
			return err
		}
	}
	writer.Flush()

	return writer.Error()

}
