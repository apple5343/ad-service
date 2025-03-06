package postgres

import (
	"server/internal/config"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"postgres-client",
		fx.Provide(
			config.NewPostgresConfig,
			NewPostgresClient,
		),
	)
}
