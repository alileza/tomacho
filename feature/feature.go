package feature

import (
	"fmt"
	"os"
	"tomato/resource"

	"gopkg.in/yaml.v2"
)

// Features
type Feature struct {
	Scenarios []Scenario `yaml:"scenarios"`
}

// Scenarios
type Scenario struct {
	ID    string `yaml:"id"`
	Steps []Step `yaml:"steps"`
}

// Scenarios
type Step struct {
	ID        string             `yaml:"id"`
	Resource  string             `yaml:"resource"`
	Action    string             `yaml:"action"`
	Arguments resource.Arguments `yaml:"arguments"`
}

func Retrieve(path string) (*Feature, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve feature %s: %w", path, err)
	}
	defer f.Close()

	var feature Feature
	if err := yaml.NewDecoder(f).Decode(&feature); err != nil {
		return nil, fmt.Errorf("Failed to decode feature %s: %w", path, err)
	}
	return &feature, nil
}
