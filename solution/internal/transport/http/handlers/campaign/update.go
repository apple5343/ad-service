package campaign

import (
	"server/internal/transport/http/handlers/campaign/converter"
	req "server/internal/transport/http/handlers/campaign/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		advertiserID := c.Param("advertiserId")
		campaignID := c.Param("campaignId")
		var campaign req.Campaign
		if err := c.Bind(&campaign); err != nil {
			return err
		}
		if err := c.Validate(campaign); err != nil {
			return err
		}
		campaignSer := converter.FromReqToCampaign(&campaign)
		campaignSer.AdvertiserID = advertiserID
		campaignSer.ID = campaignID
		campaignSer, err := h.campaignService.Update(c.Request().Context(), campaignSer)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.ToRespFromCampaign(campaignSer))
	}
}
