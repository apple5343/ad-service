package tests

import "math"

type Client struct {
	ClientID string `json:"client_id"`
	Login    string `json:"login"`
	Age      int    `json:"age"`
	Location string `json:"location"`
	Gender   string `json:"gender"`
}

type Advertiser struct {
	AdvertiserID string `json:"advertiser_id"`
	Name         string `json:"name"`
}

var genders = []string{"MALE", "FEMALE"}
var gendersAll = []string{"MALE", "FEMALE", "ALL"}

type MlScore struct {
	ClientID     string `json:"client_id"`
	AdvertiserID string `json:"advertiser_id"`
	Score        int    `json:"score"`
}

type Campaign struct {
	CampaignID        string   `json:"campaign_id,omitempty"`
	AdvertiserID      string   `json:"advertiser_id,omitempty"`
	ImpressionsLimit  *int     `json:"impressions_limit,omitempty"`
	ClicksLimit       *int     `json:"clicks_limit,omitempty"`
	CostPerClick      *float32 `json:"cost_per_click,omitempty"`
	CostPerImpression *float32 `json:"cost_per_impression,omitempty"`
	AdTitle           *string  `json:"ad_title,omitempty"`
	AdText            *string  `json:"ad_text,omitempty"`
	StartDate         *int     `json:"start_date,omitempty"`
	EndDate           *int     `json:"end_date,omitempty"`
	Target            *Target  `json:"targeting,omitempty"`
}

type Target struct {
	Gender   *string `json:"gender,omitempty"`
	AgeFrom  *int    `json:"age_from,omitempty"`
	AgeTo    *int    `json:"age_to,omitempty"`
	Location *string `json:"location,omitempty"`
}

type Ad struct {
	AdId         string `json:"ad_id"`
	AdTitle      string `json:"ad_title"`
	AdText       string `json:"ad_text"`
	AdvertiserId string `json:"advertiser_id"`
}

type Stat struct {
	ImpressionCount  int     `json:"impressions_count"`
	ClickCount       int     `json:"clicks_count"`
	Conversion       float64 `json:"conversion"`
	SpentImpressions float32 `json:"spent_impressions"`
	SpentClicks      float32 `json:"spent_clicks"`
	SpentTotal       float32 `json:"spent_total"`
}

func (s *Stat) Update() {
	s.SpentTotal = float32(math.Round(float64(s.SpentImpressions+s.SpentClicks)*100) / 100)
	if s.ImpressionCount == 0 {
		s.Conversion = 0
		return
	}
	conversion := float64(s.ClickCount) / float64(s.ImpressionCount) * 100
	s.Conversion = conversion
}

type CampaignClear struct {
	CampaignID        string  `json:"campaign_id"`
	AdvertiserID      string  `json:"advertiser_id"`
	ImpressionsLimit  int     `json:"impressions_limit"`
	ClicksLimit       int     `json:"clicks_limit"`
	CostPerClick      float32 `json:"cost_per_click"`
	CostPerImpression float32 `json:"cost_per_impression"`
	AdTitle           string  `json:"ad_title"`
	AdText            string  `json:"ad_text"`
	StartDate         int     `json:"start_date"`
	EndDate           int     `json:"end_date"`
	Target            Target  `json:"targeting"`
}
