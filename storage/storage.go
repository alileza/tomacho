package storage

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Storage struct {
	db   map[string]string
	dump map[string]string
}

func New() *Storage {
	return &Storage{make(map[string]string), make(map[string]string)}
}

func (s *Storage) Get(id, key string) (string, bool) {
	a, b := s.db[id+key]
	return a, b
}

func (s *Storage) GetMap(id, key string) (map[string]string, bool) {
	result := make(map[string]string)
	ss, ok := s.Get(id, key)
	if !ok {
		return nil, !ok
	}

	log.Println(ss)

	return result, false
}

func (s *Storage) Set(id, key string, value string) {
	s.db[id+key] = value
}

func (s *Storage) Dump(id, key string, value string) {
	s.dump[id+key] = value
}

func (s *Storage) Print() error {
	return yaml.NewEncoder(os.Stdout).Encode(s)
}
