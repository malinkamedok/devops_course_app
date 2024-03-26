package config

import (
	"github.com/caarlos0/env/v7"
	"log"
)

type Config struct {
	AppPort  string `env:"PORT" envDefault:"8000"`
	VSApiKey string `env:"API_KEY"`
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
