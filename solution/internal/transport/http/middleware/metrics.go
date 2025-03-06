package middleware

import (
	metric "server/pkg/metrics"
	"time"

	"github.com/labstack/echo/v4"
)

func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		metric.IncRequestCounter()

		start := time.Now()
		err := next(c)
		diff := time.Since(start)

		response := c.Response()
		code := response.Status

		var status string

		if code >= 200 && code < 300 {
			status = "success"
		} else {
			status = "error"
		}

		metric.IncResponseCounter(status, c.Request().Method)
		metric.HistogramResponseTimeObserve(status, diff.Seconds())
		return err
	}
}
