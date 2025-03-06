package campaign

import (
	"server/internal/repository"
	"server/internal/service"
)

type campaignService struct {
	repository repository.CampaignRepository
	objectRepo repository.CampaignObjectRepository
	ai         service.AiService
	moderate   bool
}

func NewCampaignService(repository repository.CampaignRepository, objectRepo repository.CampaignObjectRepository, ai service.AiService) service.CampaignService {
	return &campaignService{
		ai:         ai,
		repository: repository,
		objectRepo: objectRepo,
		moderate:   false,
	}
}
