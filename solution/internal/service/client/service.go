package client

import (
	"server/internal/repository"
	"server/internal/service"
)

type clientService struct {
	repo repository.ClientRepository
}

func NewClientService(repo repository.ClientRepository) service.ClientService {
	return &clientService{
		repo: repo,
	}
}
