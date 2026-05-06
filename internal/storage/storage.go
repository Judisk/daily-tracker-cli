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

type Record struct {
	Date   string
	Mood   int
	Energy int
	Focus  int
}

func NewRecord(mood, energy, focus int) (Record, error) {
	if err := validateRecord(mood, energy, focus); err != nil {
		return Record{}, err
	}

	return Record{
		Date:   time.Now().Format("2006-01-02"),
		Mood:   mood,
		Energy: energy,
		Focus:  focus,
	}, nil
}

func Save(r Record) {
	file, _ := os.OpenFile("data/data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		r.Date,
		strconv.Itoa(r.Mood),
		strconv.Itoa(r.Energy),
		strconv.Itoa(r.Focus),
	})
}

func Load() []Record {
	file, err := os.Open("data/data.csv")
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil
	}

	var records []Record

	for _, row := range rows {
		if len(row) < 4 {
			continue
		}
		mood, err1 := strconv.Atoi(row[1])
		energy, err2 := strconv.Atoi(row[2])
		focus, err3 := strconv.Atoi(row[3])

		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}

		if err := validateRecord(mood, energy, focus); err != nil {
			continue
		}

		records = append(records, Record{
			Date:   row[0],
			Mood:   mood,
			Energy: energy,
			Focus:  focus,
		})
	}

	return records
}

func validateRecord(mood, energy, focus int) error {
	if mood < MinValue || mood > MaxValue {
		return fmt.Errorf("mood must be %d-%d", MinValue, MaxValue)
	}
	if energy < MinValue || energy > MaxValue {
		return fmt.Errorf("energy must be %d-%d", MinValue, MaxValue)
	}
	if focus < MinValue || focus > MaxValue {
		return fmt.Errorf("focus must be %d-%d", MinValue, MaxValue)
	}
	return nil
}
