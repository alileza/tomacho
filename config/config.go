package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config
type Config struct {
	FeaturesPath []string   `yaml:"features_path"`
	Resources    []Resource `yaml:"resources"`
}

// Resources
type Resource struct {
	Name    string            `yaml:"name"`
	Type    string            `yaml:"type"`
	Options map[string]string `yaml:"options"`
}

func Retrieve(configPath string) (*Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open config %s: %w", configPath, err)
	}
	defer f.Close()

	var conf Config
	if err := yaml.NewDecoder(f).Decode(&conf); err != nil {
		return nil, fmt.Errorf("Failed to decode config %s: %w", configPath, err)
	}
	return &conf, nil
}
