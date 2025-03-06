package model

type Ad struct {
	ID           string `json:"ad_id"`
	Title        string `json:"ad_title"`
	Text         string `json:"ad_text"`
	AdvertiserID string `json:"advertiser_id"`
}
