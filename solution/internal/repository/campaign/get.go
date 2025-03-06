package campaign

import (
	"context"
	"server/internal/model"
	"server/internal/repository/campaign/converter"
	repo "server/internal/repository/campaign/model"

	"gorm.io/gorm"
)

func (r *campaignRepository) Get(ctx context.Context, advertiserID, campaignID string) (*model.Campaign, error) {
	campaign := &repo.Campaign{}
	result := r.db.WithContext(ctx).First(campaign, "advertiser_id = ? AND campaign_id = ?", advertiserID, campaignID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrCampaignNotFound
		}
		return nil, result.Error
	}
	return converter.FromRepoToCampaign(campaign), nil
}

func (r *campaignRepository) GetByCampaignID(ctx context.Context, campaignID string) (*model.Campaign, error) {
	campaign := &repo.Campaign{}
	result := r.db.WithContext(ctx).First(campaign, "campaign_id = ?", campaignID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrCampaignNotFound
		}
		return nil, result.Error
	}
	return converter.FromRepoToCampaign(campaign), nil
}

func (r *campaignRepository) List(ctx context.Context, advertiserID string, page, size int) ([]*model.Campaign, error) {
	var exists bool
	result := r.db.WithContext(ctx).Raw("SELECT EXISTS (SELECT 1 FROM advertisers WHERE advertiser_id = ?)", advertiserID).Scan(&exists)
	if result.Error != nil {
		return nil, result.Error
	}
	if !exists {
		return nil, ErrAdvertiserNotRegistered
	}
	campaigns := []*repo.Campaign{}
	result = r.db.WithContext(ctx).Offset((page-1)*size).Limit(size).Order("created_at DESC").Find(&campaigns, "advertiser_id = ?", advertiserID)
	if result.Error != nil {
		return nil, result.Error
	}
	return converter.FromRepoToCampaigns(campaigns), nil
}
