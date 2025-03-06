package model

type Client struct {
	ClientID *string `json:"client_id" validate:"required"`
	Login    *string `json:"login" validate:"required"`
	Age      *int    `json:"age" validate:"required"`
	Location *string `json:"location" validate:"required"`
	Gender   *string `json:"gender" validate:"required"`
}
