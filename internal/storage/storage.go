package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date   string
	Mood   int
	Energy int
	Focus  int
}

func NewRecord(mood, energy, focus int) Record {
	return Record{
		Date:   time.Now().Format("2006-01-02"),
		Mood:   mood,
		Energy: energy,
		Focus:  focus,
	}
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
	rows, _ := reader.ReadAll()

	var records []Record

	for _, row := range rows {
		mood, _ := strconv.Atoi(row[1])
		energy, _ := strconv.Atoi(row[2])
		focus, _ := strconv.Atoi(row[3])

		records = append(records, Record{
			Date:   row[0],
			Mood:   mood,
			Energy: energy,
			Focus:  focus,
		})
	}

	return records
}
