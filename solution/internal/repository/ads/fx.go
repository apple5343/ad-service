package ads

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"ads-repository",
		fx.Provide(
			NewAdsRepository,
		),
	)
}
