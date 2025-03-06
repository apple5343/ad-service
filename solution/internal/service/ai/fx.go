package ai

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"ai-service",
		fx.Provide(
			NewAiService,
		),
	)
}
