package config

import (
	"backend/models"
	"log"
)

func SyncDatabase() {
	var allModels = []interface{}{
		&models.User{},
		&models.Message{},
		&models.Patient{},
	}

	err := DB.AutoMigrate(allModels...)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

}
