package campaign

import (
	"server/internal/transport/http/handlers/campaign/converter"
	req "server/internal/transport/http/handlers/campaign/model"
	"server/pkg/errors/validate"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		advertiserID := c.Param("advertiserId")
		if _, err := uuid.Parse(advertiserID); err != nil {
			return validate.NewValidationError("invalid advertiser id")
		}

		var campaign req.Campaign
		if err := c.Bind(&campaign); err != nil {
			return err
		}
		if err := c.Validate(campaign); err != nil {
			return err
		}
		campaign.AdvertiserID = advertiserID
		campaignSer := converter.FromReqToCampaign(&campaign)
		campaignSer, err := h.campaignService.Create(c.Request().Context(), campaignSer)
		if err != nil {
			return err
		}
		return c.JSON(201, converter.ToRespFromCampaign(campaignSer))
	}
}

func (h *Handler) SaveImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		image, err := converter.ParseImage(c)
		if err != nil {
			return err
		}
		campaignID := c.Param("campaignId")
		advertiserID := c.Param("advertiserId")
		if _, err := uuid.Parse(campaignID); err != nil {
			return validate.NewValidationError("invalid campaign id")
		}
		if _, err := uuid.Parse(advertiserID); err != nil {
			return validate.NewValidationError("invalid advertiser id")
		}
		image, err = h.campaignService.SaveImage(c.Request().Context(), advertiserID, campaignID, image)
		if err != nil {
			return err
		}
		return c.JSON(200, map[string]string{"image_url": image.URL})
	}
}
