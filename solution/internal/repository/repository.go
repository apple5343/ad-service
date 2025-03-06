package repository

import (
	"context"
	"server/internal/model"
)

type ClientRepository interface {
	Save(ctx context.Context, clients []*model.Client) error
	Get(ctx context.Context, id string) (*model.Client, error)
}

type AdvertiserRepository interface {
	Save(ctx context.Context, advertisers []*model.Advertiser) error
	Get(ctx context.Context, id string) (*model.Advertiser, error)
	AddScore(ctx context.Context, score *model.Score) error
}

type CampaignRepository interface {
	Create(ctx context.Context, campaign *model.Campaign) (*model.Campaign, error)
	Get(ctx context.Context, advertiserID, campaignID string) (*model.Campaign, error)
	GetByCampaignID(ctx context.Context, campaignID string) (*model.Campaign, error)
	SaveImageUrl(ctx context.Context, imageUrl string, campaignID string) (string, error)
	Delete(ctx context.Context, advertiserID, campaignID string) error
	Update(ctx context.Context, campaign *model.Campaign) (*model.Campaign, error)
	List(ctx context.Context, advertiserID string, page, size int) ([]*model.Campaign, error)
}

type CampaignObjectRepository interface {
	SaveImage(ctx context.Context, image *model.Image) (*model.Image, error)
}

type AdsRepsitory interface {
	Get(ctx context.Context, clientID string) (string, int, float64, error)
	Impression(ctx context.Context, campaignID, clientID string, cost float64) (*model.Campaign, error)
	IsShownToClient(ctx context.Context, campaignID, clientID string) (bool, error)
	Click(ctx context.Context, campaignID, clientID string) error
}

type StatRepository interface {
	GetByCampaign(ctx context.Context, campaignID string) (*model.Stat, error)
	GetByCampaignDaily(ctx context.Context, campaignID string) (*model.Stat, error)
	GetByAdvertiser(ctx context.Context, advertiserID string) (*model.Stat, error)
	GetByAdvertiserDaily(ctx context.Context, advertiserID string) (*model.Stat, error)
}

type TimeRepository interface {
	Get(ctx context.Context) (int, error)
	Set(ctx context.Context, day int) error
	UpdateCampaignsState(ctx context.Context) error
}
