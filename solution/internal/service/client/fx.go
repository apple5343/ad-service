package client

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"service-client",
		fx.Provide(
			NewClientService,
		),
	)
}
