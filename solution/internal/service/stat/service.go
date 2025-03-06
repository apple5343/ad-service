package stat

import (
	"server/internal/repository"
	"server/internal/service"
)

type statService struct {
	statRepo repository.StatRepository
}

func NewStatService(statRepo repository.StatRepository) service.StatService {
	return &statService{
		statRepo: statRepo,
	}
}
