package s3

import (
	"server/internal/config"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"s3",
		fx.Provide(
			fx.Annotate(
				config.NewS3Config,
				fx.As(new(config.S3Config)),
			),
			NewS3Cliet,
		),
	)
}
