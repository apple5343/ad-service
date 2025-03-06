package advertiser

import "github.com/labstack/echo/v4"

func (h *Handler) Generate() echo.HandlerFunc {
	type request struct {
		AdTitle *string `json:"ad_title" validate:"required"`
	}
	return func(c echo.Context) error {
		advertiserID := c.Param("advertiserId")
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}
		description, err := h.advertiserService.Generate(c.Request().Context(), advertiserID, *req.AdTitle)
		if err != nil {
			return err
		}
		return c.JSON(200, map[string]string{"ad_text": description})
	}
}
