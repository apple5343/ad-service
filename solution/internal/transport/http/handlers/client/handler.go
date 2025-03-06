package client

import "server/internal/service"

type Handler struct {
	clientService service.ClientService
}

func NewHandler(clientService service.ClientService) *Handler {
	return &Handler{
		clientService: clientService,
	}
}
