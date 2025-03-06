package model

import (
	"context"
	"fmt"
	"server/pkg/errors/validate"
)

const (
	minUserNameLength = 1
	maxUserNameLength = 100
	minUserAge        = 0
)

type Client struct {
	ID       string
	Login    string
	Age      int
	Location string
	Gender   string
}

func (c *Client) BeforeCreate(ctx context.Context) error {
	conditions := []validate.Condition{c.ValidateAge(), c.ValidateGender(), validate.ValidateUUID(c.ID)}
	if err := validate.Validate(ctx, conditions...); err != nil {
		return err
	}
	return nil
}

func (c *Client) ValidateAge() validate.Condition {
	return func(_ context.Context) error {
		if c.Age < minUserAge {
			return validate.NewValidationError(fmt.Sprintf("age must be greater than %d", minUserAge))
		}
		return nil
	}
}

func (c *Client) ValidateLogin() validate.Condition {
	return func(_ context.Context) error {
		l := len([]rune(c.Login))
		if l < minUserNameLength || l > maxUserNameLength {
			return validate.NewValidationError(fmt.Sprintf("login length must be between %d and %d", minUserNameLength, maxUserNameLength))
		}
		return nil
	}
}

func (c *Client) ValidateGender() validate.Condition {
	return func(_ context.Context) error {
		if c.Gender != "MALE" && c.Gender != "FEMALE" {
			return validate.NewValidationError("invalid gender")
		}
		return nil
	}
}
