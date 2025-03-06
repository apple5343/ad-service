package advertiser

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"advertiser-service",
		fx.Provide(
			NewAdvertiserService,
		),
	)
}
