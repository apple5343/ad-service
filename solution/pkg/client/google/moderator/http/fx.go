package moderatorhttp

import (
	"server/internal/config"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"moderator-http",
		fx.Provide(
			config.NewPerspectiveConfig,
			NewClient,
		),
	)
}
