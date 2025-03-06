package ads

import "server/internal/service"

type Handler struct {
	adsService service.AdsService
}

func NewHandler(adsService service.AdsService) *Handler {
	return &Handler{
		adsService: adsService,
	}
}
