package storage

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Storage struct {
	DB map[string]string
}

func New() *Storage {
	return &Storage{make(map[string]string)}
}

func (s *Storage) Get(id, key string) (string, bool) {
	a, b := s.DB[id+"#"+key]
	return a, b
}

func (s *Storage) Set(id, key string, value string) {
	s.DB[id+"#"+key] = value
}

func (s *Storage) Print() error {

	return yaml.NewEncoder(os.Stdout).Encode(s)
}
