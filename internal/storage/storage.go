package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Judisk/daily-tracker-cli/internal/model"
)

func Save(r model.Record) error {

	records, err := Load()
	if err != nil {
		return fmt.Errorf("load records: %w", err)
	}

	records = append(records, r)

	data, err := json.MarshalIndent(records, "", " ")
	if err != nil {
		return fmt.Errorf("marshal records: %w", err)
	}

	path, err := dataFilePath()
	if err != nil {
		return fmt.Errorf("get data file path: %w", err)
	}

	return writeFileAtomic(path, data, 0644)
}

func Load() ([]model.Record, error) {
	path, err := dataFilePath()
	if err != nil {
		return nil, fmt.Errorf("get data file path: %w", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Record{}, nil
		}
		return nil, fmt.Errorf("read data file: %w", err)
	}

	var records []model.Record

	err = json.Unmarshal(data, &records)
	if err != nil {
		return nil, fmt.Errorf("unmarshal records: %w", err)
	}

	return records, nil

}
