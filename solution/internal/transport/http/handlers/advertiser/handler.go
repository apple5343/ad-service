package advertiser

import "server/internal/service"

type Handler struct {
	advertiserService service.AdvertiserService
}

func NewHandler(advertiserService service.AdvertiserService) *Handler {
	return &Handler{
		advertiserService: advertiserService,
	}
}
