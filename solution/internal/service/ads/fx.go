package ads

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"ads-service",
		fx.Provide(
			NewAdsService,
		),
	)
}
