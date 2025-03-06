package campaign

import (
	"context"
	"server/internal/model"
	"server/pkg/errors"
	"server/pkg/errors/validate"
)

func (s *campaignService) Create(ctx context.Context, campaign *model.Campaign) (*model.Campaign, error) {
	if err := campaign.BeforeCreate(ctx); err != nil {
		return nil, err
	}
	if s.moderate {
		if err := s.Moderate(ctx, campaign); err != nil {
			return nil, err
		}
	}
	campaign, err := s.repository.Create(ctx, campaign)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s *campaignService) SaveImage(ctx context.Context, advertiserID, campaignID string, image *model.Image) (*model.Image, error) {
	if err := validate.IsValidUUID(advertiserID); err != nil {
		return nil, validate.NewValidationError("invalid advertiser id")
	}
	if err := validate.IsValidUUID(campaignID); err != nil {
		return nil, validate.NewValidationError("invalid campaign id")
	}
	if err := image.BeforeCreate(ctx); err != nil {
		return nil, err
	}
	if _, err := s.repository.Get(ctx, advertiserID, campaignID); err != nil {
		return nil, err
	}
	if nil == image.Data {
		return nil, errors.NewError("image data is empty", errors.BadRequest)
	}
	path := "images/" + campaignID
	switch image.Type {
	case "image/jpeg":
		path += ".jpeg"
		break
	case "image/png":
		path += ".png"
		break
	case "image/jpg":
		path += ".jpg"
	default:
		return nil, errors.NewError("invalid image type", errors.BadRequest)
	}
	image.Path = path
	image, err := s.objectRepo.SaveImage(ctx, image)
	if err != nil {
		return nil, err
	}
	_, err = s.repository.SaveImageUrl(ctx, image.URL, campaignID)
	if err != nil {
		return nil, err
	}
	return image, nil
}
