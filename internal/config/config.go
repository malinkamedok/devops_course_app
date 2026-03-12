package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	AppPort  string   `env:"PORT" envDefault:"8000"`
	ApiKeys  []string `env:"API_KEYS"`
	ChatID   string   `env:"CHAT_ID"`
	ApiToken string   `env:"API_TOKEN"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Println("Error in parsing env")
		return nil, err
	}
	if len(cfg.ApiKeys) == 0 {
		return nil, fmt.Errorf("no API_KEYS provided")
	}
	return cfg, nil
}
