package config

import "github.com/ilyakaznacheev/cleanenv"

type RedisConfig interface {
	Address() string
	Password() string
	DB() int
}

type redisConfig struct {
	AddressField  string `env:"REDIS_ADDRESS" env-default:"localhost:6379"`
	PasswordField string `env:"REDIS_PASSWORD" env-default:""`
	DBField       int    `env:"REDIS_DB" env-default:"0"`
}

func NewRedisConfig() (RedisConfig, error) {
	cfg := redisConfig{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *redisConfig) Address() string {
	return c.AddressField
}

func (c *redisConfig) Password() string {
	return c.PasswordField
}

func (c *redisConfig) DB() int {
	return c.DBField
}
