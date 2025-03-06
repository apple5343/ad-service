package ads

import (
	"server/internal/transport/http/converter"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Click() echo.HandlerFunc {
	type request struct {
		ClientID *string `json:"client_id" validate:"required"`
	}
	return func(c echo.Context) error {
		campaignID := c.Param("adId")
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}
		err := h.adsService.Click(c.Request().Context(), campaignID, converter.FromStringPtr(req.ClientID))
		if err != nil {
			return err
		}
		return c.JSON(204, nil)
	}
}
