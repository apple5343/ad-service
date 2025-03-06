package campaign_object_mock

import (
	"context"
	"server/internal/model"
	"server/internal/repository"

	"github.com/brianvoe/gofakeit"
	"go.uber.org/fx"
)

type campaignObjectRepository struct {
}

func NewCampaignObjectRepository() repository.CampaignObjectRepository {
	return &campaignObjectRepository{}
}

func (r *campaignObjectRepository) SaveImage(ctx context.Context, image *model.Image) (*model.Image, error) {
	image.URL = gofakeit.ImageURL(100, 100)
	return image, nil
}

func NewModule() fx.Option {
	return fx.Module(
		"campaign-object-mock-repository",
		fx.Provide(NewCampaignObjectRepository),
	)
}
