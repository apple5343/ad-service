package model

type Client struct {
	ClientID string `gorm:"primaryKey;type:uuid"`
	Login    string `gorm:"not null"`
	Age      int    `gorm:"not null"`
	Location string `gorm:"not null"`
	Gender   string `gorm:"not null"`
}
