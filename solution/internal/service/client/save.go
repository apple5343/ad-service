package client

import (
	"context"
	"server/internal/model"
)

func (s *clientService) Save(ctx context.Context, clients []*model.Client) error {
	for _, client := range clients {
		if err := client.BeforeCreate(ctx); err != nil {
			return err
		}
	}
	return s.repo.Save(ctx, clients)
}
