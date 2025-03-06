package advertiser

import (
	"server/internal/repository"
	"server/internal/service"
)

type advertiserService struct {
	repo      repository.AdvertiserRepository
	aiService service.AiService
}

func NewAdvertiserService(repo repository.AdvertiserRepository, aiService service.AiService) service.AdvertiserService {
	return &advertiserService{
		repo:      repo,
		aiService: aiService,
	}
}
