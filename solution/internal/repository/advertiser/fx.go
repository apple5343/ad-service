package advertiser

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"advertiser-repository",
		fx.Provide(
			NewAdvertiserRepository,
		),
	)
}
