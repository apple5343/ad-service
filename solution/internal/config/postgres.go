package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type PostgresConfig interface {
	DSN() string
}

type postgresConfig struct {
	DSNField string `env:"POSTGRES_DSN" env-default:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
}

func NewPostgresConfig() (PostgresConfig, error) {
	cfg := postgresConfig{}
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println("POSTGRES_DSN:", cfg.DSNField)
	return &cfg, nil
}

func (p *postgresConfig) DSN() string {
	return p.DSNField
}
