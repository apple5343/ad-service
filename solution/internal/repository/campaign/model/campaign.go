package model

import (
	"database/sql"
	"time"
)

type Campaign struct {
	CampaignID        string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	AdvertiserID      string  `gorm:"not null"`
	AdTitle           string  `gorm:"not null"`
	AdText            string  `gorm:"not null"`
	ClicksLimit       int     `gorm:"not null"`
	ImpressionsLimit  int     `gorm:"not null"`
	CostPerClick      float64 `gorm:"not null"`
	CostPerImpression float64 `gorm:"not null"`
	StartDate         int     `gorm:"not null"`
	EndDate           int     `gorm:"not null"`
	Active            bool    `gorm:"not null"`
	ImageUrl          sql.NullString
	Target            CampaignTarget `gorm:"embedded"`
	CreatedAt         time.Time      `gorm:"not null"`
}

type CampaignTarget struct {
	Gender   sql.NullString
	AgeFrom  sql.NullInt64
	AgeTo    sql.NullInt64
	Location sql.NullString
}
