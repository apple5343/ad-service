package metric

import (
	"server/internal/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	requestCounter        prometheus.Counter
	responseCounter       *prometheus.CounterVec
	histogramResponseTime *prometheus.HistogramVec
	adActionCounter       *prometheus.CounterVec
}

var myPrometheus *Prometheus

func (p *Prometheus) initMetrics(cfg config.PrometheusConfig) (*Metrics, error) {
	m := &Metrics{
		requestCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: cfg.Space(),
				Subsystem: "http",
				Name:      cfg.Name() + "_requests_total",
				Help:      "Количество запросов к серверу",
			},
		),
		responseCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: cfg.Space(),
				Subsystem: "http",
				Name:      cfg.Name() + "_responses_total",
				Help:      "Количество ответов от сервера",
			},
			[]string{"status", "method"},
		),
		histogramResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: cfg.Space(),
				Subsystem: "http",
				Name:      cfg.Name() + "_histogram_response_time_seconds",
				Help:      "Время ответа от сервера",
				Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
			},
			[]string{"status"},
		),
		adActionCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: cfg.Space(),
				Subsystem: "http",
				Name:      cfg.Name() + "_ad_actions_total",
				Help:      "Действия по камапаниям",
			},
			[]string{"action"},
		),
	}

	return m, nil
}

func IncRequestCounter() {
	if myPrometheus == nil {
		return
	}
	myPrometheus.metrics.requestCounter.Inc()
}

func IncResponseCounter(status string, method string) {
	if myPrometheus == nil {
		return
	}
	myPrometheus.metrics.responseCounter.WithLabelValues(status, method).Inc()
}

func HistogramResponseTimeObserve(status string, time float64) {
	if myPrometheus == nil {
		return
	}
	myPrometheus.metrics.histogramResponseTime.WithLabelValues(status).Observe(time)
}

func IncAdActionCounter(action string) {
	if myPrometheus == nil {
		return
	}
	myPrometheus.metrics.adActionCounter.WithLabelValues(action).Inc()
}
