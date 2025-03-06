package converter

import (
	"server/internal/model"
	req "server/internal/transport/http/handlers/ads/model"
)

func FromCampaignToAd(campaign *model.Campaign) *req.Ad {
	return &req.Ad{
		ID:           campaign.ID,
		Title:        campaign.AdTitle,
		Text:         campaign.AdText,
		AdvertiserID: campaign.AdvertiserID,
	}
}
