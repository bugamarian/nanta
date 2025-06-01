package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	NotesDir string `yaml:"notes_dir"`
	Savemode string `yaml:"savemode"` // "timestamp" or "daily"
	Editor   string `yaml:"editor"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}
