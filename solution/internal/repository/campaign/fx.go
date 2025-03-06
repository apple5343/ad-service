package campaign

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"campaign-repository",
		fx.Provide(
			NewCampaignRepository,
		),
	)
}
