package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	MinValue = 0
	MaxValue = 5
)
const (
	PillsMin          = 0
	PillsMax          = 50
	PillsLowThreshold = 7
)

type Record struct {
	Date   string
	Mood   int
	Energy int
	Focus  int
	Pills  int
}

func NewRecord(mood, energy, focus, pills int) (Record, error) {

	r := Record{
		Date:   time.Now().Format("2006-01-02"),
		Mood:   mood,
		Energy: energy,
		Focus:  focus,
		Pills:  pills,
	}

	if err := r.Validate(); err != nil {
		return Record{}, err
	}
	return r, nil
}

func Save(r Record) error {
	file, err := os.OpenFile("data/data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		if err := writer.Write([]string{"Date", "Mood", "Energy", "Focus", "Pills"}); err != nil {
			return err
		}
	}

	if err := writer.Write([]string{
		r.Date,
		strconv.Itoa(r.Mood),
		strconv.Itoa(r.Energy),
		strconv.Itoa(r.Focus),
		strconv.Itoa(r.Pills),
	}); err != nil {
		return err
	}
	writer.Flush()

	return writer.Error()

}

func Load() ([]Record, error) {
	file, err := os.Open("data/data.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return []Record{}, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []Record

	for _, row := range rows {
		if len(row) < 5 {
			continue
		}
		mood, err1 := strconv.Atoi(row[1])
		energy, err2 := strconv.Atoi(row[2])
		focus, err3 := strconv.Atoi(row[3])
		pills, err4 := strconv.Atoi(row[4])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			continue
		}

		r := Record{
			Date:   row[0],
			Mood:   mood,
			Energy: energy,
			Focus:  focus,
			Pills:  pills,
		}

		if err := r.Validate(); err != nil {
			continue
		}

		records = append(records, r)

	}

	return records, nil
}

func (r Record) Validate() error {
	if r.Mood < MinValue || r.Mood > MaxValue {
		return fmt.Errorf("mood must be %d-%d", MinValue, MaxValue)
	}
	if r.Energy < MinValue || r.Energy > MaxValue {
		return fmt.Errorf("energy must be %d-%d", MinValue, MaxValue)
	}
	if r.Focus < MinValue || r.Focus > MaxValue {
		return fmt.Errorf("focus must be %d-%d", MinValue, MaxValue)
	}
	if r.Pills < PillsMin || r.Pills > PillsMax {
		return fmt.Errorf("pills must be %d-%d", PillsMin, PillsMax)
	}
	return nil
}
