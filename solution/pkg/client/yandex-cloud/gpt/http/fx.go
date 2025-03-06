package gpthttp

import (
	"server/internal/config"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"gpt-http",
		fx.Provide(
			config.NewYandexGptConfig,
			NewClient,
		),
	)
}
