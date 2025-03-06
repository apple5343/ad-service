package campaign

import (
	"server/internal/transport/http/handlers/campaign/converter"
	"server/pkg/errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		campaignID := c.Param("campaignId")
		advertiserID := c.Param("advertiserId")
		campaign, err := h.campaignService.Get(c.Request().Context(), advertiserID, campaignID)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.ToRespFromCampaign(campaign))
	}
}

func (h *Handler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		advertiserID := c.Param("advertiserId")
		page := c.Request().URL.Query().Get("page")
		size := c.Request().URL.Query().Get("size")
		sizeInt, err := strconv.Atoi(size)
		if err != nil {
			if size == "" {
				sizeInt = 5
			} else {
				return errors.NewError("invalid size", errors.BadRequest)
			}
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			if page == "" {
				pageInt = 1
			} else {
				return errors.NewError("invalid page", errors.BadRequest)
			}
		}
		campaigns, err := h.campaignService.List(c.Request().Context(), advertiserID, pageInt, sizeInt)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.FromCampaignsToResp(campaigns))
	}
}
