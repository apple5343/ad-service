package model

type Score struct {
	ClientID     *string `json:"client_id" validate:"required"`
	AdvertiserID *string `json:"advertiser_id" validate:"required"`
	Score        *int    `json:"score" validate:"required"`
}
