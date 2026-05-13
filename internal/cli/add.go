package cli

import (
	"bufio"
	"fmt"

	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/storage"
)

func RunAdd(r *bufio.Reader) error {

	fields := NewFields()
	record, err := NewRecord(r, fields)
	if err != nil {
		return err
	}

	if record.Pills <= model.PillsLowThreshold {
		fmt.Printf("Warning pills running low (%d left)\n", record.Pills)
	}

	if err := storage.Save(record); err != nil {

		return fmt.Errorf("Error saving data:%w", err)
	}
	fmt.Println("Saved ✅")
	return nil
}
