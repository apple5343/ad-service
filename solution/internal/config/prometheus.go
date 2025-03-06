package config

import "github.com/ilyakaznacheev/cleanenv"

type PrometheusConfig interface {
	Space() string
	Name() string
	Address() string
}

type prometheusConfig struct {
	SpaceField string `env:"PROMETHEUS_SPACE" env-default:"default"`
	NameField  string `env:"PROMETHEUS_NAME" env-default:"prometheus"`
	TargetAddr string `env:"PROMETHEUS_TARGET_ADDR" env-default:":2112"`
}

func NewPrometheusConfig() (PrometheusConfig, error) {
	cfg := prometheusConfig{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (p *prometheusConfig) Space() string {
	return p.SpaceField
}

func (p *prometheusConfig) Name() string {
	return p.NameField
}

func (p *prometheusConfig) Address() string {
	return p.TargetAddr
}
