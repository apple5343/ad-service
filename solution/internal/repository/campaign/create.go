package campaign

import (
	"context"
	"database/sql"
	"server/internal/model"
	"server/internal/repository/campaign/converter"
	repo "server/internal/repository/campaign/model"
	"server/pkg/client/postgres"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *campaignRepository) Create(ctx context.Context, campaign *model.Campaign) (*model.Campaign, error) {
	campaignRepo := converter.ToRepoFromCampaign(campaign)
	result := r.db.WithContext(ctx).Create(campaignRepo)
	if result.Error != nil {
		if pqErr, ok := result.Error.(*pgconn.PgError); ok {
			if pqErr.Code == postgres.CodeForeignKeyViolation {
				return nil, ErrAdvertiserNotRegistered
			}
		}
		return nil, result.Error
	}
	return converter.FromRepoToCampaign(campaignRepo), nil
}

func (r *campaignRepository) SaveImageUrl(ctx context.Context, imageUrl string, campaignID string) (string, error) {
	url := sql.NullString{}
	if imageUrl != "" {
		url = sql.NullString{
			String: imageUrl,
			Valid:  true,
		}
	}
	res := r.db.WithContext(ctx).Model(&repo.Campaign{CampaignID: campaignID}).Where("campaign_id = ?", campaignID).Update("image_url", url)
	if res.Error != nil {
		if res.Error == sql.ErrNoRows {
			return "", ErrCampaignNotFound
		}
		return "", res.Error
	}
	if res.RowsAffected == 0 {
		return "", ErrCampaignNotFound
	}
	return imageUrl, nil
}
