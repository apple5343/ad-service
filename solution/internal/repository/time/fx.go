package time

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"repository-time",
		fx.Provide(
			NewTimeRepository,
		),
	)
}
