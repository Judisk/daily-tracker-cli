package cli

import (
	"bufio"
	"fmt"

	application "github.com/Judisk/daily-tracker-cli/internal/app"
	"github.com/Judisk/daily-tracker-cli/internal/model"
)

func Add(r *bufio.Reader, service *application.Service) error {
	fields := newFields()

	record, err := newRecord(r, fields)
	if err != nil {
		return err
	}

	if record.Pills <= model.PillsLowThreshold {
		fmt.Printf("Warning pills running low (%d left)\n", record.Pills)
	}

	if err := service.CreateRecord(record); err != nil {
		return fmt.Errorf("create record: %w", err)
	}
	fmt.Println("Saved ✅")
	return nil
}
