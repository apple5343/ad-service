package model

import (
	"context"
	"fmt"
	"server/pkg/errors/validate"
)

const (
	minAdvertiserNameLength = 1
	maxAdvertiserNameLength = 100
)

type Advertiser struct {
	ID   string
	Name string
}

func (a *Advertiser) BeforeCreate(ctx context.Context) error {
	conditions := []validate.Condition{a.ValidateName(), validate.ValidateUUID(a.ID)}
	return validate.Validate(ctx, conditions...)
}

func (a *Advertiser) ValidateName() validate.Condition {
	return func(_ context.Context) error {
		l := len([]rune(a.Name))
		if l < minAdvertiserNameLength || l > maxAdvertiserNameLength {
			return validate.NewValidationError(fmt.Sprintf("advertiser name length must be between %d and %d", minAdvertiserNameLength, maxAdvertiserNameLength))
		}
		return nil
	}
}
