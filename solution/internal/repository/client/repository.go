package client

import (
	"server/internal/repository"

	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) repository.ClientRepository {
	return &clientRepository{
		db: db,
	}
}
