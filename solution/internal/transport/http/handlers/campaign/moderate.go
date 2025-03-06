package campaign

import "github.com/labstack/echo/v4"

func (h *Handler) Moderate() echo.HandlerFunc {
	type request struct {
		Enabled bool `json:"enabled" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}
		err := h.campaignService.SetModerateStatus(c.Request().Context(), req.Enabled)
		if err != nil {
			return err
		}
		return c.JSON(200, map[string]bool{"enabled": req.Enabled})
	}
}
