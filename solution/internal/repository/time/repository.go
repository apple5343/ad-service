package time

import (
	"server/internal/repository"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type timeRepository struct {
	rdb *redis.Client
	db  *gorm.DB
}

func NewTimeRepository(rdb *redis.Client, db *gorm.DB) repository.TimeRepository {
	return &timeRepository{
		rdb: rdb,
		db:  db,
	}
}
