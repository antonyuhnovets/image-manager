package config

import (
	"errors"
	"os"
)

type Config struct {
	AMQP_URL     string
	STORAGE      string
	STORAGE_PATH string
}

func LoadConfig() (*Config, error) {
	storage := os.Getenv("STORAGE")
	if storage == "" {
		storage = "local"
	}
	path := os.Getenv("STORAGE_PATH")
	if path == "" {
		path = "./test-images"
	}
	amqp := os.Getenv("AMQP_URL")
	if amqp == "" {
		return nil, errors.New("AMQP connection string not setted")
	}
	cfg := &Config{
		amqp,
		storage,
		path,
	}

	return cfg, nil
}
