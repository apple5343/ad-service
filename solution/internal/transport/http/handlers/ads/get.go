package ads

import (
	"server/internal/transport/http/handlers/ads/converter"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID := c.Request().URL.Query().Get("client_id")
		campaign, err := h.adsService.Get(c.Request().Context(), clientID)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.FromCampaignToAd(campaign))
	}
}
