package client

import (
	"context"
	"server/internal/model"
	"server/pkg/errors/validate"
)

func (s *clientService) Get(ctx context.Context, id string) (*model.Client, error) {
	if err := validate.IsValidUUID(id); err != nil {
		return nil, err
	}
	return s.repo.Get(ctx, id)
}
