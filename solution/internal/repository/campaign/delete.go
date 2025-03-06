package campaign

import (
	"context"
	repo "server/internal/repository/campaign/model"

	"gorm.io/gorm"
)

func (r *campaignRepository) Delete(ctx context.Context, advertiserID, campaignID string) error {
	res := r.db.WithContext(ctx).Delete(&repo.Campaign{}, "campaign_id = ? AND advertiser_id = ?", campaignID, advertiserID)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return ErrCampaignNotFound
		}
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrCampaignNotFound
	}
	return nil
}
