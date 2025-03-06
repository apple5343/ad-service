package campaign

import "server/internal/service"

type Handler struct {
	campaignService service.CampaignService
}

func NewHandler(campaignService service.CampaignService) *Handler {
	return &Handler{
		campaignService: campaignService,
	}
}
