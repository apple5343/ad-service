package campaign

import (
	"context"
	"server/internal/model"
	"server/internal/repository/campaign/converter"

	"gorm.io/gorm"
)

func (r *campaignRepository) Update(ctx context.Context, campaign *model.Campaign) (*model.Campaign, error) {
	campaignRepo := converter.ToRepoFromCampaign(campaign)
	result := r.db.WithContext(ctx).Save(campaignRepo)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrCampaignNotFound
		}
		return nil, result.Error
	}

	return converter.FromRepoToCampaign(campaignRepo), nil
}
