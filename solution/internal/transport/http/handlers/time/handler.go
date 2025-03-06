package time

import "server/internal/service"

type Handler struct {
	timeService service.TimeService
}

func NewHandler(timeService service.TimeService) *Handler {
	return &Handler{
		timeService: timeService,
	}
}
