package converter

import (
	"server/internal/model"

	repo "server/internal/repository/advertiser/model"
)

func ToRepoFromAdvertiser(advertiser *model.Advertiser) *repo.Advertiser {
	return &repo.Advertiser{
		AdvertiserID: advertiser.ID,
		Name:         advertiser.Name,
	}
}

func ToAdvertiserFromRepo(advertiser *repo.Advertiser) *model.Advertiser {
	return &model.Advertiser{
		ID:   advertiser.AdvertiserID,
		Name: advertiser.Name,
	}
}

func ToAdvertisersFromRepo(advertisers []*repo.Advertiser) []*model.Advertiser {
	ads := make([]*model.Advertiser, len(advertisers))
	for i, advertiser := range advertisers {
		ads[i] = ToAdvertiserFromRepo(advertiser)
	}
	return ads
}

func ToRepoFromAdvertisers(advertisers []*model.Advertiser) []*repo.Advertiser {
	ads := make([]*repo.Advertiser, len(advertisers))
	for i, advertiser := range advertisers {
		ads[i] = ToRepoFromAdvertiser(advertiser)
	}
	return ads
}
