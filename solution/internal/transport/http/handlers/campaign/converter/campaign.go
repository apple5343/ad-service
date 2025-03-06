package converter

import (
	"server/internal/model"
	"server/internal/transport/http/converter"
	req "server/internal/transport/http/handlers/campaign/model"
)

func FromReqToCampaign(campaign *req.Campaign) *model.Campaign {
	return &model.Campaign{
		AdvertiserID:      campaign.AdvertiserID,
		AdTitle:           converter.FromStringPtr(campaign.AdTitle),
		AdText:            converter.FromStringPtr(campaign.AdText),
		ClicksLimit:       converter.FromIntPtr(campaign.ClicksLimit),
		ImpressionsLimit:  converter.FromIntPtr(campaign.ImpressionsLimit),
		CostPerClick:      converter.FromFloat64Ptr(campaign.CostPerClick),
		CostPerImpression: converter.FromFloat64Ptr(campaign.CostPerImpression),
		StartDate:         converter.FromIntPtr(campaign.StartDate),
		EndDate:           converter.FromIntPtr(campaign.EndDate),
		Target:            FromReqToCampaignTarget(campaign.Target),
	}
}

func FromReqToCampaignTarget(target req.CampaignTarget) model.CampaignTarget {
	return model.CampaignTarget{
		Gender:   converter.ToNullString(target.Gender),
		AgeFrom:  converter.ToNullInt64(target.AgeFrom),
		AgeTo:    converter.ToNullInt64(target.AgeTo),
		Location: converter.ToNullString(target.Location),
	}
}

func ToRespFromCampaign(campaign *model.Campaign) *req.Campaign {
	return &req.Campaign{
		CampaignID:        campaign.ID,
		AdvertiserID:      campaign.AdvertiserID,
		ImpressionsLimit:  &campaign.ImpressionsLimit,
		ClicksLimit:       &campaign.ClicksLimit,
		CostPerClick:      &campaign.CostPerClick,
		CostPerImpression: &campaign.CostPerImpression,
		AdTitle:           &campaign.AdTitle,
		AdText:            &campaign.AdText,
		StartDate:         &campaign.StartDate,
		EndDate:           &campaign.EndDate,
		ImageUrl:          campaign.ImageUrl,
		Target:            ToRespFromCampaignTarget(campaign.Target),
	}
}

func ToRespFromCampaignTarget(target model.CampaignTarget) req.CampaignTarget {
	return req.CampaignTarget{
		Gender:   converter.ToString(target.Gender),
		AgeFrom:  converter.ToInt(target.AgeFrom),
		AgeTo:    converter.ToInt(target.AgeTo),
		Location: converter.ToString(target.Location),
	}
}

func FromCampaignsToResp(campaigns []*model.Campaign) []*req.Campaign {
	resp := make([]*req.Campaign, 0, len(campaigns))
	for _, campaign := range campaigns {
		resp = append(resp, ToRespFromCampaign(campaign))
	}
	return resp
}
