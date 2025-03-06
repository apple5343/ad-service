package s3

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"campaign-object-s3-repository",
		fx.Provide(NewCampaignObjectRepository),
	)
}
