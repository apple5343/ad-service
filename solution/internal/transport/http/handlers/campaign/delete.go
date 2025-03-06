package campaign

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		advertiserID := c.Param("advertiserId")
		campaignID := c.Param("campaignId")
		err := h.campaignService.Delete(c.Request().Context(), advertiserID, campaignID)
		if err != nil {
			return err
		}
		return c.JSON(204, nil)
	}
}
