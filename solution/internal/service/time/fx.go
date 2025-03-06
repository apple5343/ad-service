package time

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"time-service",
		fx.Provide(
			NewTimeService,
		),
	)
}
