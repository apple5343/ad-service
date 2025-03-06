package ads

import (
	"server/internal/repository"
	"server/internal/service"
)

type adsService struct {
	adsRepsitory repository.AdsRepsitory
}

func NewAdsService(adsRepsitory repository.AdsRepsitory) service.AdsService {
	return &adsService{
		adsRepsitory: adsRepsitory,
	}
}
