package storage

import (
	"encoding/json"
	"os"

	"github.com/Judisk/daily-tracker-cli/internal/model"
)

func Save(r model.Record) error {

	records, err := Load()
	if err != nil {
		return err
	}

	records = append(records, r)

	data, err := json.MarshalIndent(records, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/data.json", data, 0644)
}

func Load() ([]model.Record, error) {
	data, err := os.ReadFile("data/data.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Record{}, nil
		}
		return nil, err
	}

	var records []model.Record

	err = json.Unmarshal(data, &records)
	if err != nil {
		return nil, err
	}

	return records, nil

}
