package main

import (
	"backend/config"
	"backend/handlers"
	"backend/models"
	"backend/routes"
	"time"

	"github.com/google/uuid"

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

	// Register AI routes
	routes.SetupAIRoutes(r)

	// Start handling messages
	go handlers.HandleMessages()

	// Start server
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

}

func createFixtures() {
	now := time.Now()
	patients := []models.Patient{
		{
			UUID:      uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "grey",
			Phone:     "+33601234567",
			Keywords: []string{
				"keyword1",
				"keyword2",
			},
		},
		{
			UUID:      uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "grey",
			Phone:     "+33701234599",
			Keywords: []string{
				"keyword3",
				"keyword4",
			},
		},
		// Ajoutez d'autres patients si nécessaire
	}

	for _, patient := range patients {
		config.DB.Create(&patient)
	}
}
