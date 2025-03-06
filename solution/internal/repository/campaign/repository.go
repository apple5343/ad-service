package campaign

import (
	"server/internal/repository"

	"gorm.io/gorm"
)

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) repository.CampaignRepository {
	return &campaignRepository{
		db: db,
	}
}
