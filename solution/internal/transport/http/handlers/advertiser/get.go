package advertiser

import (
	"server/internal/transport/http/handlers/advertiser/converter"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("advertiserId")
		advertiser, err := h.advertiserService.Get(c.Request().Context(), id)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.ToRespFromAdvertiser(advertiser))
	}
}
