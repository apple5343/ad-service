package config

import "github.com/ilyakaznacheev/cleanenv"

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	AddressField string `env:"HTTP_ADDRESS" env-default:"localhost:8080"`
}

func NewHTTPConfig() (HTTPConfig, error) {
	cfg := httpConfig{}
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *httpConfig) Address() string {
	return c.AddressField
}
