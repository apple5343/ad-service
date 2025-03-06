package service

import (
	"context"
	"server/internal/model"
)

type ClientService interface {
	Save(ctx context.Context, clients []*model.Client) error
	Get(ctx context.Context, id string) (*model.Client, error)
}

type AdvertiserService interface {
	Save(ctx context.Context, advertisers []*model.Advertiser) error
	Get(ctx context.Context, id string) (*model.Advertiser, error)
	AddScore(ctx context.Context, score *model.Score) error
	Generate(ctx context.Context, adveristerID, campaignName string) (string, error)
}

type CampaignService interface {
	Create(ctx context.Context, campaign *model.Campaign) (*model.Campaign, error)
	SaveImage(ctx context.Context, advertiserID, campaignID string, image *model.Image) (*model.Image, error)
	Get(ctx context.Context, advertiserID, campaignID string) (*model.Campaign, error)
	Delete(ctx context.Context, advertiserID, campaignID string) error
	Update(ctx context.Context, campaign *model.Campaign) (*model.Campaign, error)
	List(ctx context.Context, advertiserID string, page, size int) ([]*model.Campaign, error)
	SetModerateStatus(ctx context.Context, moderate bool) error
}

type AdsService interface {
	Get(ctx context.Context, clientID string) (*model.Campaign, error)
	Click(ctx context.Context, campaignID, clientID string) error
}

type StatService interface {
	GetByCampaign(ctx context.Context, campaignID string) (*model.Stat, error)
	GetByCampaignDaily(ctx context.Context, campaignID string) (*model.Stat, error)
	GetByAdvertiser(ctx context.Context, advertiserID string) (*model.Stat, error)
	GetByAdvertiserDaily(ctx context.Context, advertiserID string) (*model.Stat, error)
}

type AiService interface {
	ModerateText(ctx context.Context, text string) (bool, []string, error)
	GenerateCampaignDescription(ctx context.Context, adveristerName, campaignName string) (string, error)
}

type TimeService interface {
	Set(ctx context.Context, day int) error
}
