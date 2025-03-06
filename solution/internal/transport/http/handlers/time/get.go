package time

import (
	"server/pkg/time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		date := time.Day()
		return c.JSON(200, map[string]int{"current_date": date})
	}
}
