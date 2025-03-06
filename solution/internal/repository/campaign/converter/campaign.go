package converter

import (
	"database/sql"
	"server/internal/model"
	repo "server/internal/repository/campaign/model"
)

func ToRepoFromCampaign(campaign *model.Campaign) *repo.Campaign {
	image := sql.NullString{}
	if campaign.ImageUrl != "" {
		image = sql.NullString{
			String: campaign.ImageUrl,
			Valid:  true,
		}
	}
	return &repo.Campaign{
		CampaignID:        campaign.ID,
		AdvertiserID:      campaign.AdvertiserID,
		AdTitle:           campaign.AdTitle,
		AdText:            campaign.AdText,
		ClicksLimit:       campaign.ClicksLimit,
		ImpressionsLimit:  campaign.ImpressionsLimit,
		CostPerClick:      campaign.CostPerClick,
		CostPerImpression: campaign.CostPerImpression,
		StartDate:         campaign.StartDate,
		EndDate:           campaign.EndDate,
		ImageUrl:          image,
		Active:            campaign.Active,
		Target:            ToRepoFromCampaignTarget(campaign.Target),
	}
}

func FromRepoToCampaign(campaign *repo.Campaign) *model.Campaign {
	return &model.Campaign{
		ID:                campaign.CampaignID,
		AdvertiserID:      campaign.AdvertiserID,
		AdTitle:           campaign.AdTitle,
		AdText:            campaign.AdText,
		ClicksLimit:       campaign.ClicksLimit,
		ImpressionsLimit:  campaign.ImpressionsLimit,
		CostPerClick:      campaign.CostPerClick,
		CostPerImpression: campaign.CostPerImpression,
		StartDate:         campaign.StartDate,
		EndDate:           campaign.EndDate,
		ImageUrl:          campaign.ImageUrl.String,
		Active:            campaign.Active,
		Target:            FromRepoToCampaignTarget(campaign.Target),
	}
}

func ToRepoFromCampaignTarget(target model.CampaignTarget) repo.CampaignTarget {
	return repo.CampaignTarget{
		Gender:   target.Gender,
		AgeFrom:  target.AgeFrom,
		AgeTo:    target.AgeTo,
		Location: target.Location,
	}
}

func FromRepoToCampaignTarget(target repo.CampaignTarget) model.CampaignTarget {
	return model.CampaignTarget{
		Gender:   target.Gender,
		AgeFrom:  target.AgeFrom,
		AgeTo:    target.AgeTo,
		Location: target.Location,
	}
}

func FromRepoToCampaigns(campaigns []*repo.Campaign) []*model.Campaign {
	var result []*model.Campaign
	for _, campaign := range campaigns {
		result = append(result, FromRepoToCampaign(campaign))
	}
	return result
}
