package config

import "github.com/ilyakaznacheev/cleanenv"

type PerspectiveConfig interface {
	URL() string
}

type perspectiveConfig struct {
	Endpoint string `env:"PERSPECTIVE_ENDPOINT" env-default:"https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze"`
	Key      string `env:"PERSPECTIVE_KEY"`
}

func NewPerspectiveConfig() (PerspectiveConfig, error) {
	cfg := perspectiveConfig{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *perspectiveConfig) URL() string {
	return c.Endpoint + "?key=" + c.Key
}
