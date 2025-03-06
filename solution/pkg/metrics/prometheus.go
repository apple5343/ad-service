package metric

import (
	"context"
	"log"
	"net/http"
	"server/internal/config"
	"server/pkg/logger"

	"go.uber.org/fx"
)

type Prometheus struct {
	metrics *Metrics
	server  *http.Server
	cfg     config.PrometheusConfig
}

func NewPrometheus(cfg config.PrometheusConfig) (*Prometheus, error) {
	p := &Prometheus{cfg: cfg}
	m, err := p.initMetrics(cfg)
	if err != nil {
		return nil, err
	}
	p.metrics = m
	p.initServer()
	myPrometheus = p
	return p, nil
}

func NewModule() fx.Option {
	return fx.Module(
		"metrics",
		fx.Provide(
			config.NewPrometheusConfig,
			NewPrometheus,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, p *Prometheus) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						p.initServer()
						go func() {
							logger.Info("starting prometheus server")
							if err := p.runServer(); err != nil {
								log.Fatal(err)
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return p.server.Shutdown(ctx)
					},
				})
			},
		),
	)
}
