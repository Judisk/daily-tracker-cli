package storage

import (
	"os"
	"path/filepath"
)

var getDataFilePath = dataFilePath

func dataFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, "health-tracker-go")

	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(appDir, "data.json"), nil
}
