package config

import (
	"encoding/json"
	"log"
	"os"
)

type Object struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Symlink    string `json:"symlink"`
	Background string `json:"background"`
	Foreground string `json:"foreground"`
	Icon       string `json:"icon"`
}

type MetaConfig struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type Config struct {
	EnableTls  bool        `json:"enableTls"`
	Avatar     *string     `json:"avatar"`
	Name       *string     `json:"name"`
	Bio        *string     `json:"bio"`
	Objects    []Object    `json:"objects"`
	Background *string     `json:"background"`
	Meta       *MetaConfig `json:"meta"`
}

func LoadConfig(path string) *Config {
	conf, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config
	err = json.Unmarshal(conf, &cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal json: %v", err)
	}

	return &cfg
}
