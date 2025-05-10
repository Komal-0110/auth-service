package db

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabse() {
	cfg := config.AppConfig

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Failed to auto-migrate: %v", err)
	}

	log.Println("✅ Database connected!")
	DB = database
}
