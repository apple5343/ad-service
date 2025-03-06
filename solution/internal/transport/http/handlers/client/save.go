package client

import (
	"server/internal/transport/http/handlers/client/converter"
	"server/internal/transport/http/handlers/client/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Save() echo.HandlerFunc {
	return func(c echo.Context) error {
		var clients []*model.Client
		if err := c.Bind(&clients); err != nil {
			return err
		}
		for _, client := range clients {
			if err := c.Validate(client); err != nil {
				return err
			}
		}
		err := h.clientService.Save(c.Request().Context(), converter.FromReqToClients(clients))
		if err != nil {
			return err
		}
		return c.JSON(201, clients)
	}
}
