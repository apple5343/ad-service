package advertiser

import (
	"context"
	"server/internal/model"
)

func (s *advertiserService) Save(ctx context.Context, advertisers []*model.Advertiser) error {
	for _, advertiser := range advertisers {
		if err := advertiser.BeforeCreate(ctx); err != nil {
			return err
		}
	}
	return s.repo.Save(ctx, advertisers)
}
