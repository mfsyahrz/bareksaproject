package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Service  Service
	Postgres Postgres
}

type Service struct {
	Name string `env:"SERVICE_NAME"`
	Port Port
}

type Port struct {
	REST string `env:"SERVICE_PORT_REST"`
}

type Postgres struct {
	User            string `env:"POSTGRES_USER,required"`
	Password        string `env:"POSTGRES_PASSWORD,required"`
	Name            string `env:"POSTGRES_NAME,required"`
	Port            string `env:"POSTGRES_PORT,default=5432"`
	Host            string `env:"POSTGRES_HOST,default=localhost"`
	MaxOpenConns    int    `env:"POSTGRES_MAX_OPEN_CONNS,default=5"`
	MaxConnLifetime int    `env:"POSTGRES_MAX_CONN,default=10"`
	MaxIdleLifetime int    `env:"POSTGRES_MAX_IDLE,default=5"`
}

func New(envFile string) (*Config, error) {

	_ = godotenv.Load(envFile)

	config := Config{}
	if err := envdecode.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
