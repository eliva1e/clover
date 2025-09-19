package main

import (
	"encoding/json"
	"log"
	"os"
)

type Link struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Symlink string `json:"symlink"`
}

type Config struct {
	Address string `json:"address"`
	Avatar  string `json:"avatar"`
	Name    string `json:"name"`
	Bio     string `json:"bio"`
	Links   []Link `json:"links"`
}

func LoadConfig() Config {
	conf, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("failed to read config: " + err.Error())
	}

	var cfg Config
	err = json.Unmarshal(conf, &cfg)
	if err != nil {
		log.Fatal("failed to unmarshal json: " + err.Error())
	}

	return cfg
}
