package advertiser

import (
	"server/internal/transport/http/handlers/advertiser/converter"
	"server/internal/transport/http/handlers/advertiser/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Save() echo.HandlerFunc {
	return func(c echo.Context) error {
		var advertisers []*model.Advertiser
		if err := c.Bind(&advertisers); err != nil {
			return err
		}
		for _, advertiser := range advertisers {
			if err := c.Validate(advertiser); err != nil {
				return err
			}
		}
		err := h.advertiserService.Save(c.Request().Context(), converter.ToAdvertisersFromReq(advertisers))
		if err != nil {
			return err
		}
		return c.JSON(201, advertisers)
	}
}
