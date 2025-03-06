package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (p *Prometheus) initServer() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	prometheusServer := &http.Server{
		Addr:    p.cfg.Address(),
		Handler: mux,
	}
	p.server = prometheusServer
}

func (p *Prometheus) runServer() error {
	return p.server.ListenAndServe()
}
