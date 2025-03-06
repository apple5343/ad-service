package postgres

import (
	"server/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient(cfg config.PostgresConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(40)
	sqlDB.SetMaxOpenConns(200)
	return db, nil
}
