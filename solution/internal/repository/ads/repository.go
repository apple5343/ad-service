package ads

import (
	"server/internal/repository"

	"gorm.io/gorm"
)

type adsRepsitory struct {
	db *gorm.DB
}

func NewAdsRepository(db *gorm.DB) repository.AdsRepsitory {
	return &adsRepsitory{
		db: db,
	}
}
