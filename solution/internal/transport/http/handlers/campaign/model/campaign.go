package model

type Campaign struct {
	CampaignID        string         `json:"campaign_id"`
	AdvertiserID      string         `json:"advertiser_id"`
	ImpressionsLimit  *int           `json:"impressions_limit" validate:"required"`
	ClicksLimit       *int           `json:"clicks_limit" validate:"required"`
	CostPerClick      *float64       `json:"cost_per_click" validate:"required"`
	CostPerImpression *float64       `json:"cost_per_impression" validate:"required"`
	AdTitle           *string        `json:"ad_title" validate:"required"`
	AdText            *string        `json:"ad_text" validate:"required"`
	StartDate         *int           `json:"start_date" validate:"required"`
	EndDate           *int           `json:"end_date" validate:"required"`
	ImageUrl          string         `json:"image_url,omitempty"`
	Target            CampaignTarget `json:"targeting"`
}

type CampaignTarget struct {
	Location *string `json:"location,omitempty"`
	AgeFrom  *int    `json:"age_from,omitempty"`
	AgeTo    *int    `json:"age_to,omitempty"`
	Gender   *string `json:"gender,omitempty"`
}
