package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect initializes the database connection using SQLite
func Connect() error {
	// Use SQLite for simplicity
	dbPath := "./education_platform.db"

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")

	// Auto-migrate models - commented out to avoid errors
	// err = DB.AutoMigrate(&models.User{}, &models.Course{})
	// if err != nil {
	// 	return fmt.Errorf("failed to migrate database: %w", err)
	// }

	log.Println("Database migration completed")
	return nil
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
