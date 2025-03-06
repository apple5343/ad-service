package advertiser

import (
	"context"
	"server/internal/model"
	"server/pkg/errors/validate"
)

func (s *advertiserService) Get(ctx context.Context, id string) (*model.Advertiser, error) {
	if err := validate.IsValidUUID(id); err != nil {
		return nil, err
	}
	return s.repo.Get(ctx, id)
}
