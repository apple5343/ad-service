package advertiser

import (
	"context"
	"server/internal/model"
	"server/internal/repository/advertiser/converter"
	repo "server/internal/repository/advertiser/model"

	"gorm.io/gorm"
)

func (s *advertiserRepository) Get(ctx context.Context, id string) (*model.Advertiser, error) {
	advertiser := &repo.Advertiser{}
	result := s.db.WithContext(ctx).First(advertiser, "advertiser_id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrAdvertiserNotFound
		}
		return nil, result.Error
	}
	return converter.ToAdvertiserFromRepo(advertiser), nil
}
