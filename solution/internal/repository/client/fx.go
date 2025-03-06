package client

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"repository-client",
		fx.Provide(
			NewClientRepository,
		),
	)
}
