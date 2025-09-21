package config

import (
	"encoding/json"
	"log"
	"os"
)

type Link struct {
	IsLabel    bool   `json:"isLabel"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Symlink    string `json:"symlink"`
	Background string `json:"background"`
	Foreground string `json:"foreground"`
}

type Config struct {
	Address string `json:"address"`
	Avatar  string `json:"avatar"`
	Name    string `json:"name"`
	Bio     string `json:"bio"`
	Links   []Link `json:"links"`
}

func LoadConfig(path string) Config {
	conf, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config
	err = json.Unmarshal(conf, &cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal json: %v", err)
	}

	return cfg
}
