package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Save(root string, cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")

	if err != nil {
		return fmt.Errorf(
			"failed to encode config: %w",
			err,
		)
	}

	data = append(data, '\n')

	path := filepath.Join(root, ConfigFileName)

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf(
			"failed to write %s: %w",
			ConfigFileName,
			err,
		)
	}

	return nil
}
