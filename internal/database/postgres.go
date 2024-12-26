package database

import (
	"fmt"
	"go-rest-api-kafka/internal/config"
	"go-rest-api-kafka/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

		// TODO: скорее всего надо убрать
    // Автомиграция таблиц
    err = db.AutoMigrate(&models.Plan{})
    if err != nil {
        return nil, err
    }

    return db, nil
}