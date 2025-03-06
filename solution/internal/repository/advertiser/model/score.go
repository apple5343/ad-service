package model

type Score struct {
	AdvertiserID string `gorm:"primaryKey;type:uuid;not null;references:advertisers(advertiser_id)"`
	ClientID     string `gorm:"primaryKey;type:uuid;not null;references:clients(client_id)"`
	Score        int    `gorm:"not null"`
}
