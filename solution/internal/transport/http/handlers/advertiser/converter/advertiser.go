package converter

import (
	"server/internal/transport/http/converter"
	req "server/internal/transport/http/handlers/advertiser/model"

	"server/internal/model"
)

func ToAdvertiserFromReq(advertiser *req.Advertiser) *model.Advertiser {
	return &model.Advertiser{
		ID:   converter.FromStringPtr(advertiser.AdvertiserID),
		Name: converter.FromStringPtr(advertiser.Name),
	}
}

func ToRespFromAdvertiser(advertiser *model.Advertiser) *req.Advertiser {
	return &req.Advertiser{
		AdvertiserID: &advertiser.ID,
		Name:         &advertiser.Name,
	}
}

func ToAdvertisersFromReq(advertisers []*req.Advertiser) []*model.Advertiser {
	ads := make([]*model.Advertiser, len(advertisers))
	for i, advertiser := range advertisers {
		ads[i] = ToAdvertiserFromReq(advertiser)
	}
	return ads
}

func ToRespFromAdvertisers(advertisers []*model.Advertiser) []*req.Advertiser {
	ads := make([]*req.Advertiser, len(advertisers))
	for i, advertiser := range advertisers {
		ads[i] = ToRespFromAdvertiser(advertiser)
	}
	return ads
}
