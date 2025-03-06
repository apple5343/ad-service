package model

type Stat struct {
	ImpressionsCount int     `json:"impressions_count"`
	ClicksCount      int     `json:"clicks_count"`
	Conversion       float64 `json:"conversion"`
	SpentImpressions float64 `json:"spent_impressions"`
	SpentClicks      float64 `json:"spent_clicks"`
	SpentTotal       float64 `json:"spent_total"`
}
