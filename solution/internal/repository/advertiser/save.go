package advertiser

import (
	"context"
	"server/internal/model"
	"server/internal/repository/advertiser/converter"

	"github.com/jackc/pgx/v5/pgconn"
)

func (s *advertiserRepository) Save(ctx context.Context, advertisers []*model.Advertiser) error {
	result := s.db.WithContext(ctx).Save(converter.ToRepoFromAdvertisers(advertisers))
	if result.Error != nil {
		if result.Error != nil {
			if pqErr, ok := result.Error.(*pgconn.PgError); ok {
				if pqErr.Code == "21000" {
					return ErrAdvertiserAlreadyExists
				}
			}
			return result.Error
		}
	}
	return nil
}
