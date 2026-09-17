package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Load(root string) (Config, error) {
	path := filepath.Join(root, ConfigFileName)

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf(
			"failed to read %s: %w",
			ConfigFileName,
			err,
		)
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf(
			"invalid %s: %w",
			ConfigFileName,
			err,
		)
	}

	return cfg, nil
}
