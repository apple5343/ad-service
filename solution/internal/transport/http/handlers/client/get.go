package client

import (
	"server/internal/transport/http/handlers/client/converter"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("clientId")
		client, err := h.clientService.Get(c.Request().Context(), id)
		if err != nil {
			return err
		}
		return c.JSON(200, converter.FromClientToResp(client))
	}
}
