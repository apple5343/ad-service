package advertiser

import (
	"server/internal/repository"

	"gorm.io/gorm"
)

type advertiserRepository struct {
	db *gorm.DB
}

func NewAdvertiserRepository(db *gorm.DB) repository.AdvertiserRepository {
	return &advertiserRepository{
		db: db,
	}
}
