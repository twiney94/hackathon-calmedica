package config

import (
	"backend/models"
	"log"
)

func SyncDatabase() {
	var allModels = []interface{}{
		&models.User{},
	}

	err := DB.AutoMigrate(allModels...)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

}
