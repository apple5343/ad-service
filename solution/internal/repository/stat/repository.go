package stat

import (
	"server/internal/repository"

	"gorm.io/gorm"
)

type statRepository struct {
	db *gorm.DB
}

func NewStatRepository(db *gorm.DB) repository.StatRepository {
	return &statRepository{
		db: db,
	}
}
