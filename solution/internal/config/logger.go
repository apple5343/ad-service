package config

import "github.com/ilyakaznacheev/cleanenv"

type LoggerConfig interface {
	Level() string
}

type loggerConfig struct {
	LevelField string `env:"LOG_LEVEL" env-default:"debug"`
}

func NewLoggerConfig() (LoggerConfig, error) {
	var cfg loggerConfig
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (l *loggerConfig) Level() string {
	return l.LevelField
}
