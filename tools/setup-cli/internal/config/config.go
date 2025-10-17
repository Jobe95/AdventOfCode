package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".aoc-config"

// GetConfigPath returns the path to the config file in the user's home directory
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, configFileName), nil
}

// SetSession saves the AOC session cookie
func SetSession(session string) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, []byte(session), 0600)
}

// GetSession retrieves the AOC session cookie
func GetSession() (string, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // No config file yet
		}
		return "", fmt.Errorf("failed to read config: %w", err)
	}

	return string(data), nil
}
