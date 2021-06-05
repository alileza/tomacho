package feature

import (
	"fmt"
	"os"
	"tomato/resource"

	"github.com/cucumber/godog/colors"
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

// Step
type Step struct {
	Resource  string             `yaml:"resource"`
	Action    string             `yaml:"action"`
	Arguments resource.Arguments `yaml:"arguments"`
}

func (s Step) String() string {
	return fmt.Sprintf("%s: %s\t=> %v", colors.Bold(colors.White)(s.Resource), s.Action, s.Arguments)
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
