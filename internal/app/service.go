package app

import (
	"fmt"

	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/storage"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateRecord(record model.Record) error {
	if err := storage.Save(record); err != nil {
		return fmt.Errorf("save record: %w", err)
	}
	return nil
}
