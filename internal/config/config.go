package config

import (
	"log"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	AppPort string   `env:"PORT" envDefault:"8000"`
	ApiKeys []string `env:"API_KEYS"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Println("Error in parsing env")
		return nil, err
	}
	return cfg, nil
}
