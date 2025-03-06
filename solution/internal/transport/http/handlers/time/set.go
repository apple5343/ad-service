package time

import (
	"server/pkg/time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Set() echo.HandlerFunc {
	type Request struct {
		CurrentDate int `json:"current_date" validate:"required"`
	}
	return func(c echo.Context) error {
		var req Request
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			req.CurrentDate = time.Day() + 1
		}

		err := h.timeService.Set(c.Request().Context(), req.CurrentDate)
		if err != nil {
			return err
		}
		return c.JSON(200, req)
	}
}
