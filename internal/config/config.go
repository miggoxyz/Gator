// Package config manages the application configuration.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// Config represents the application configuration.
type Config struct {
	DBUrl string `json:"db_url"`
	User  string `json:"current_user_name"`
}

// SetUser updates the current user and writes the config to file.
func (cfg *Config) SetUser(username string) error {
	cfg.User = username
	return write(*cfg)
}

// Read loads the configuration from the file.
func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get config file path: %w", err)
	}
	content, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return Config{}, nil
	} else if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}
	var cfg Config
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return cfg, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get config file path: %w", err)
	}
	encoding, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()
	_, err = file.Write(encoding)
	if err != nil {
		return fmt.Errorf("failed to write config to file: %w", err)
	}
	return nil
}

// getConfigFilePath returns the full path to the config file.
func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homedir, configFileName), nil
}