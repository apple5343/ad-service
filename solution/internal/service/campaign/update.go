package campaign

import (
	"context"
	"server/internal/model"
	"server/pkg/errors"
	"server/pkg/errors/validate"

	"github.com/google/uuid"
)

func (s *campaignService) Update(ctx context.Context, campaign *model.Campaign) (*model.Campaign, error) {
	if _, err := uuid.Parse(campaign.ID); err != nil {
		return nil, validate.NewValidationError("invalid campaign id")
	}
	if _, err := uuid.Parse(campaign.AdvertiserID); err != nil {
		return nil, validate.NewValidationError("invalid advertiser id")
	}
	if err := campaign.BeforeUpdate(ctx); err != nil {
		return nil, err
	}
	campaignRepo, err := s.repository.GetByCampaignID(ctx, campaign.ID)
	if err != nil {
		return nil, err
	}
	if campaignRepo.AdvertiserID != campaign.AdvertiserID {
		return nil, errors.NewError("permission denied", errors.Forbidden)
	}

	if s.moderate {
		if err := s.Moderate(ctx, campaign); err != nil {
			return nil, err
		}
	}

	if campaignRepo.Active {
		if campaign.ImpressionsLimit != campaignRepo.ImpressionsLimit || campaign.ClicksLimit != campaignRepo.ClicksLimit ||
			campaign.StartDate != campaignRepo.StartDate || campaign.EndDate != campaignRepo.EndDate {
			return nil, errors.NewError("campaign is active", errors.BadRequest)
		}
	} else {
		campaignRepo.ImpressionsLimit = campaign.ImpressionsLimit
		campaignRepo.ClicksLimit = campaign.ClicksLimit
		campaignRepo.StartDate = campaign.StartDate
		campaignRepo.EndDate = campaign.EndDate
	}
	campaignRepo.CostPerClick = campaign.CostPerClick
	campaignRepo.CostPerImpression = campaign.CostPerImpression
	campaignRepo.AdTitle = campaign.AdTitle
	campaignRepo.AdText = campaign.AdText
	campaignRepo.Target = campaign.Target
	campaignRepo.SetActive()
	res, err := s.repository.Update(ctx, campaignRepo)
	if err != nil {
		return nil, err
	}
	return res, nil
}
