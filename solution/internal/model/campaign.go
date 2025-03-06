package model

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"server/pkg/errors"
	"server/pkg/errors/validate"
	"server/pkg/time"
	"slices"
)

const (
	minAdTitleLength = 1
	maxAdTitleLength = 255

	minAdTextLength = 1
	maxAdTextLength = 32767

	minClicksLimit     = 0
	minImpressionLimit = 0

	minCostPerClick      = 0
	minCostPerImpression = 0
)

var gender = []string{"MALE", "FEMALE", "ALL"}

type Campaign struct {
	ID                string
	AdvertiserID      string
	AdTitle           string
	AdText            string
	ClicksLimit       int
	ImpressionsLimit  int
	CostPerClick      float64
	CostPerImpression float64
	StartDate         int
	EndDate           int
	Active            bool
	ImageUrl          string
	Target            CampaignTarget
}

type CampaignTarget struct {
	Gender   sql.NullString
	AgeFrom  sql.NullInt64
	AgeTo    sql.NullInt64
	Location sql.NullString
}

type Image struct {
	URL  string
	Data io.Reader
	Type string
	Path string
}

func (i *Image) BeforeCreate(ctx context.Context) error {
	if i.Data == nil {
		return nil
	}
	switch i.Type {
	case "image/jpeg":
		i.Type = "image/jpeg"
		break
	case "image/png":
		i.Type = "image/png"
		break
	case "image/jpg":
		i.Type = "image/jpg"
	default:
		return errors.NewError("invalid image type", errors.BadRequest)
	}
	return nil
}

func (c *Campaign) BeforeCreate(ctx context.Context) error {
	conditions := []validate.Condition{
		c.ValidateClicksLimit(), c.ValidateImpressionLimit(), c.ValidateCostPerClick(), c.ValidateCostPerImpression(),
		c.ValidateStartDate(), c.ValidateEndDate(), c.ValidateDay(false), validate.ValidateUUID(c.AdvertiserID), c.ValidateLimits()}
	if err := validate.Validate(ctx, conditions...); err != nil {
		return err
	}

	if err := c.ValidateTarget(); err != nil {
		return err
	}
	c.SetActive()
	return nil
}

func (c *Campaign) BeforeUpdate(ctx context.Context) error {
	conditions := []validate.Condition{
		c.ValidateClicksLimit(), c.ValidateImpressionLimit(), c.ValidateCostPerClick(), c.ValidateCostPerImpression(),
		c.ValidateStartDate(), c.ValidateEndDate(), c.ValidateDay(true), validate.ValidateUUID(c.AdvertiserID), c.ValidateLimits()}
	if err := validate.Validate(ctx, conditions...); err != nil {
		return err
	}

	if err := c.ValidateTarget(); err != nil {
		return err
	}
	c.SetActive()
	return nil
}

func (c *Campaign) SetActive() {
	day := time.Day()
	if day >= c.StartDate && day <= c.EndDate {
		c.Active = true
	}
}

func (c *Campaign) ValidateTarget() error {
	if c.Target.Gender.Valid && !slices.Contains(gender, c.Target.Gender.String) {
		return validate.NewValidationError("invalid gender")
	}
	if c.Target.AgeFrom.Valid && c.Target.AgeTo.Valid && c.Target.AgeFrom.Int64 > c.Target.AgeTo.Int64 {
		return validate.NewValidationError("age from must be less than age to")
	}
	if c.Target.AgeFrom.Valid && c.Target.AgeFrom.Int64 < 0 {
		return validate.NewValidationError("age from must be greater than 0")
	}
	if c.Target.AgeTo.Valid && c.Target.AgeTo.Int64 < 0 {
		return validate.NewValidationError("age to must be greater than 0")
	}
	if c.Target.Location.Valid && len([]rune(c.Target.Location.String)) == 0 {
		return validate.NewValidationError("location must not be empty")
	}
	return nil
}

func (c *Campaign) ValidateAdTitle() validate.Condition {
	return func(_ context.Context) error {
		l := len([]rune(c.AdTitle))
		if l < minAdTitleLength || l > maxAdTitleLength {
			return validate.NewValidationError(fmt.Sprintf("ad title length must be between %d and %d", minAdTitleLength, maxAdTitleLength))
		}
		return nil
	}
}

func (c *Campaign) ValidateDay(isUpdate bool) validate.Condition {
	return func(_ context.Context) error {
		day := time.Day()
		if !isUpdate && day > c.StartDate || day > c.EndDate {
			return validate.NewValidationError("invalid start date or end date")
		}
		if c.StartDate > c.EndDate {
			return validate.NewValidationError("start date must be less than end date")
		}
		return nil
	}
}

func (c *Campaign) ValidatetAdText() validate.Condition {
	return func(_ context.Context) error {
		l := len([]rune(c.AdText))
		if l < minAdTextLength || l > maxAdTextLength {
			return validate.NewValidationError(fmt.Sprintf("ad text length must be between %d and %d", minAdTextLength, maxAdTextLength))
		}
		return nil
	}
}

func (c *Campaign) ValidateClicksLimit() validate.Condition {
	return func(_ context.Context) error {
		if c.ClicksLimit < minClicksLimit {
			return validate.NewValidationError(fmt.Sprintf("clicks limit must be greater than %d", minClicksLimit))
		}
		return nil
	}
}

func (c *Campaign) ValidateImpressionLimit() validate.Condition {
	return func(_ context.Context) error {
		if c.ImpressionsLimit < minImpressionLimit {
			return validate.NewValidationError(fmt.Sprintf("impression limit must be greater than %d", minImpressionLimit))
		}
		return nil
	}
}

func (c *Campaign) ValidateCostPerClick() validate.Condition {
	return func(_ context.Context) error {
		if c.CostPerClick < minCostPerClick {
			return validate.NewValidationError(fmt.Sprintf("cost per click must be greater than %f", minCostPerClick))
		}
		return nil
	}
}

func (c *Campaign) ValidateCostPerImpression() validate.Condition {
	return func(_ context.Context) error {
		if c.CostPerImpression < minCostPerImpression {
			return validate.NewValidationError(fmt.Sprintf("cost per impression must be greater than %f", minCostPerImpression))
		}
		return nil
	}
}

func (c *Campaign) ValidateStartDate() validate.Condition {
	return func(_ context.Context) error {
		if c.StartDate < 0 {
			return validate.NewValidationError("start date must be greater than 0")
		}
		return nil
	}
}

func (c *Campaign) ValidateEndDate() validate.Condition {
	return func(_ context.Context) error {
		if c.EndDate < 0 {
			return validate.NewValidationError("end date must be greater than 0")
		}
		return nil
	}
}

func (c *Campaign) ValidateLimits() validate.Condition {
	return func(_ context.Context) error {
		if c.ClicksLimit > c.ImpressionsLimit {
			return validate.NewValidationError("clicks limit must be less than impressions limit")
		}
		return nil
	}
}
