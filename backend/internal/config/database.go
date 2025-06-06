package config

import (
	"fmt"
	"log"

	"github.com/45ai/backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.Database.DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(
		&models.User{},
		&models.Template{},
		&models.Task{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("Database connection established and migrations completed")
	return db, nil
}
