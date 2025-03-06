package converter

import (
	"server/internal/model"
	repo "server/internal/repository/campaign/model"
)

func ToCampaignFromRepo(campaign *repo.Campaign) *model.Campaign {
	return &model.Campaign{
		ID:           campaign.CampaignID,
		AdvertiserID: campaign.AdvertiserID,
		AdTitle:      campaign.AdTitle,
		AdText:       campaign.AdText,
	}
}
