package stat

import (
	"server/internal/service"
)

type Handler struct {
	statService service.StatService
}

func NewHandler(statService service.StatService) *Handler {
	return &Handler{
		statService: statService,
	}
}
