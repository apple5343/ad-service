package redis

import (
	"server/internal/config"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"redis-client",
		fx.Provide(
			config.NewRedisConfig,
			NewRedisClient,
		),
	)
}
