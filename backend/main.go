package main

import (
	"backend/config"
	"backend/handlers"
	"backend/models"
	"backend/routes"
	"github.com/google/uuid"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDb()
	config.SyncDatabase()
	config.LoadEnvVariables()
	createFixtures()
}

func main() {
	//db connection result := config.DB.Create(&user)
	db := config.DB

	r := gin.Default()

	// Register routes
	routes.SetupUserRoutes(r, db)
	routes.RegisterPatientRoutes(r)

	// Register message routes
	routes.SetupMessageRoutes(r, db)

	// Register chat routes
	routes.SetupChatRoutes(r)

	// Start handling messages
	go handlers.HandleMessages()

	// Start server
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

}

func createFixtures() {
	now := time.Now() // Capture du moment actuel pour CreatedAt et UpdatedAt
	patients := []models.Patient{
		{
			UUID:      uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "bleu",
			Phone:     "1234567897",
		},
		{
			UUID:      uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "bleu",
			Phone:     "1234543216",
		},
		// Ajoutez d'autres patients si nécessaire
	}

	for _, patient := range patients {
		config.DB.Create(&patient) // Création de l'entrée en base de données
	}
}
