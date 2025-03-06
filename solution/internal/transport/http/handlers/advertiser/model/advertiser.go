package model

type Advertiser struct {
	AdvertiserID *string `json:"advertiser_id" validate:"required"`
	Name         *string `json:"name,required" validate:"required"`
}
