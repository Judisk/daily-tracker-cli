package storage

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Judisk/daily-tracker-cli/internal/model"
)

func useTestDataFile(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")

	oldGetDataFilePath := getDataFilePath

	getDataFilePath = func() (string, error) {
		return path, nil
	}

	t.Cleanup(func() {
		getDataFilePath = oldGetDataFilePath
	})

	return path
}

func TestLoadReturnsEmptyRecordsIfFileDoesNotExist(t *testing.T) {
	useTestDataFile(t)

	records, err := Load()
	if err != nil {
		t.Fatalf("Load() unexpected error: %v", err)
	}

	if records == nil {
		t.Fatalf("Load() returned nil records, want empty slice")
	}

	if len(records) != 0 {
		t.Fatalf("Load() returned %d records, want 0", len(records))
	}
}

func TestSaveWritesRecord(t *testing.T) {
	path := useTestDataFile(t)

	record := model.Record{
		Date:         time.Now(),
		SleepQuality: 3,
		Mood:         4,
		Energy:       2,
		Focus:        5,
		Pills:        10,
		Notes:        "test note",
		SideEffects:  "none",
	}

	if err := Save(record); err != nil {
		t.Fatalf("Save() unexpected error: %v", err)
	}

	if _, err := Load(); err != nil {
		t.Fatalf("Load() after Save() unexpected error: %v", err)
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("data file was not created: %v", err)
	}
}

func TestLoadReadsSavedRecord(t *testing.T) {
	useTestDataFile(t)

	want := model.Record{
		Date:         time.Now(),
		SleepQuality: 3,
		Mood:         4,
		Energy:       2,
		Focus:        5,
		Pills:        10,
		Notes:        "test note",
		SideEffects:  "none",
	}

	if err := Save(want); err != nil {
		t.Fatalf("Save() unexpected error: %v", err)
	}

	records, err := Load()
	if err != nil {
		t.Fatalf("Load() unexpected error: %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("Load() returned %d records, want 1", len(records))
	}

	got := records[0]

	if got.SleepQuality != want.SleepQuality {
		t.Errorf("SleepQuality = %d, want %d", got.SleepQuality, want.SleepQuality)
	}

	if got.Mood != want.Mood {
		t.Errorf("Mood = %d, want %d", got.Mood, want.Mood)
	}

	if got.Energy != want.Energy {
		t.Errorf("Energy = %d, want %d", got.Energy, want.Energy)
	}

	if got.Focus != want.Focus {
		t.Errorf("Focus = %d, want %d", got.Focus, want.Focus)
	}

	if got.Pills != want.Pills {
		t.Errorf("Pills = %d, want %d", got.Pills, want.Pills)
	}

	if got.Notes != want.Notes {
		t.Errorf("Notes = %q, want %q", got.Notes, want.Notes)
	}

	if got.SideEffects != want.SideEffects {
		t.Errorf("SideEffects = %q, want %q", got.SideEffects, want.SideEffects)
	}
}

func TestLoadReturnsErrorForInvalidJSON(t *testing.T) {
	path := useTestDataFile(t)

	err := os.WriteFile(path, []byte("not json"), 0644)
	if err != nil {
		t.Fatalf("write invalid json: %v", err)
	}

	records, err := Load()
	if err == nil {
		t.Fatal("Load() expected error, got nil")
	}

	if records != nil {
		t.Fatalf("Load() records = %+v, want nil", records)
	}
}

func TestLoadReturnsErrorWhenDataPathFails(t *testing.T) {
	oldGetDataFilePath := getDataFilePath

	getDataFilePath = func() (string, error) {
		return "", errors.New("test data path error")
	}

	t.Cleanup(func() {
		getDataFilePath = oldGetDataFilePath
	})

	records, err := Load()

	if err == nil {
		t.Fatal("Load() expected error, got nil")
	}
	if records != nil {
		t.Fatalf("Load() records = %+v, want nil", records)
	}
}

func TestSaveReturnsErrorWhenDataPathFails(t *testing.T) {
	oldGetDataFilePath := getDataFilePath

	getDataFilePath = func() (string, error) {
		return "", errors.New("test data path error")
	}

	t.Cleanup(func() {
		getDataFilePath = oldGetDataFilePath
	})

	err := Save(model.Record{})

	if err == nil {
		t.Fatal("Save() expected error, got nil")
	}
}
