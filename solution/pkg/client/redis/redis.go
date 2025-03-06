package redis

import (
	"context"
	"server/internal/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password(),
		DB:       cfg.DB(),
	})
	if err := c.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return c, nil
}
