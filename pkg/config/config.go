package config

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Broker  string `env:"BROKER" env-default:"rabbitmq"`
	Url     string `env:"CONNECTION_URL"`
	Storage string `env:"STORAGE" env-default:"local"`
	Path    string `env:"STORAGE_PATH" env-default:"./test-images"`
	Host    string `env:"HOST" env-default:"0.0.0.0"`
	Port    string `env:"PORT" env-default:"8080"`
}

// Load config from enviroment
// Throw an error if broker connection string is not setted
func LoadConfig() (Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, err
	}
	if cfg.Url == "" {
		return cfg, errors.New("broker connection string not setted")
	}

	return cfg, nil
}
