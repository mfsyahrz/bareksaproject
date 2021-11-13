package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Service Service
}

type Service struct {
	Name string `env:"SERVICE_NAME"`
	Port Port
}

type Port struct {
	REST string `env:"SERVICE_PORT_REST"`
}

func New(envFile string) (*Config, error) {

	_ = godotenv.Load(envFile)

	config := Config{}
	if err := envdecode.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
