package config

import "os"

type Config struct {
	AMQP_URL     string `env:"AMQP_URL"`
	STORAGE      string `env:"STORAGE" default:"local"`
	STORAGE_PATH string `env:"STORAGE_PATH"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	cfg.AMQP_URL = os.Getenv("AMQP_URL")
	cfg.STORAGE = os.Getenv("STORAGE")
	cfg.STORAGE_PATH = os.Getenv("STORAGE_PATH")

	return cfg
}
