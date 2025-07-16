package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Remote struct {
		URL    string `yaml:"url"`
		APIKey string `yaml:"api_key"`
	} `yaml:"remote"`

	Logging struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"logging"`
}

// LoadConfig loads configuration from the specified file
func LoadConfig(filename string) (*Config, error) {
	config := &Config{}

	// Read the config file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
	}

	// Parse YAML
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", filename, err)
	}

	// Set defaults if not specified
	if config.Server.Port == "" {
		config.Server.Port = "8080"
	}
	if config.Server.Host == "" {
		config.Server.Host = "localhost"
	}
	if config.Remote.URL == "" {
		config.Remote.URL = "http://localhost:8000"
	}
	if config.Logging.Level == "" {
		config.Logging.Level = "info"
	}

	return config, nil
}

// LoadDefaultConfig loads configuration from the default config file
func LoadDefaultConfig() (*Config, error) {
	return LoadConfig("config.yaml")
}

// CreateDefaultConfig creates a default config file if it doesn't exist
func CreateDefaultConfig(filename string) error {
	// Check if file already exists
	if _, err := os.Stat(filename); err == nil {
		return nil // File already exists
	}

	config := &Config{}
	config.Server.Port = "8080"
	config.Server.Host = "localhost"
	config.Remote.URL = "http://localhost:8000"
	config.Remote.APIKey = "your-api-key-here"
	config.Logging.Level = "info"
	config.Logging.File = "logs/secauto-gui.log"

	// Marshal to YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %w", err)
	}

	// Create config directory if it doesn't exist
	dir := "config"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Write the file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
