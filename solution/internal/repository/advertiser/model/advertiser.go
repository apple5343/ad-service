package model

type Advertiser struct {
	AdvertiserID string `gorm:"primaryKey;type:uuid"`
	Name         string `gorm:"not null"`
}
