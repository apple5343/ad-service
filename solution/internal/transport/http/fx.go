package http

import (
	"context"
	"log"
	"server/internal/config"
	"server/pkg/logger"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"http-server",
		fx.Provide(
			config.NewHTTPConfig,
			NewServer,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, s *Server) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						logger.Debug("starting http server")
						go func() {
							if err := s.Run(); err != nil {
								log.Fatal(err)
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return s.Stop(ctx)
					},
				})
			},
		),
	)
}
