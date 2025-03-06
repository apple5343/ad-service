package advertiser

import (
	"server/internal/transport/http/handlers/advertiser/converter"
	req "server/internal/transport/http/handlers/advertiser/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AddScore() echo.HandlerFunc {
	return func(c echo.Context) error {
		var score req.Score
		if err := c.Bind(&score); err != nil {
			return err
		}
		if err := c.Validate(score); err != nil {
			return err
		}
		err := h.advertiserService.AddScore(c.Request().Context(), converter.ToScoreFromReq(&score))
		if err != nil {
			return err
		}
		return c.JSON(200, score)
	}
}
