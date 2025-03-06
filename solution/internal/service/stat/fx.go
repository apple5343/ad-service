package stat

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"stat-service",
		fx.Provide(
			NewStatService,
		),
	)
}
