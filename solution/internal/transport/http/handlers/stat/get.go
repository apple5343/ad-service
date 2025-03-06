package stat

import (
	"server/internal/transport/http/handlers/stat/converter"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetByCampaign() echo.HandlerFunc {
	return func(c echo.Context) error {
		campaignID := c.Param("campaignId")
		stat, err := h.statService.GetByCampaign(c.Request().Context(), campaignID)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.FromStatToResp(stat))
	}
}

func (h *Handler) GetByCampaignDaily() echo.HandlerFunc {
	return func(c echo.Context) error {
		campaignID := c.Param("campaignId")
		stat, err := h.statService.GetByCampaignDaily(c.Request().Context(), campaignID)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.FromStatToResp(stat))
	}
}

func (h *Handler) GetByAdvertiser() echo.HandlerFunc {
	return func(c echo.Context) error {
		advertiserID := c.Param("advertiserId")
		stat, err := h.statService.GetByAdvertiser(c.Request().Context(), advertiserID)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.FromStatToResp(stat))
	}
}

func (h *Handler) GetByAdvertiserDaily() echo.HandlerFunc {
	return func(c echo.Context) error {
		advertiserID := c.Param("advertiserId")
		stat, err := h.statService.GetByAdvertiserDaily(c.Request().Context(), advertiserID)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.FromStatToResp(stat))
	}
}
