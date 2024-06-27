package main

import (
	"backend/config"
	"backend/handlers"
	"backend/models"
	"backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"time"
)

func init() {
	config.ConnectToDb()
	config.SyncDatabase()
	config.LoadEnvVariables()
	createFixtures()
}

func main() {
	// db connection result := config.DB.Create(&user)
	db := config.DB

	r := gin.Default()

	// Configure CORS middleware
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Apply the CORS middleware
	r.Use(func(c *gin.Context) {
		corsConfig.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})

	// Register routes
	routes.SetupUserRoutes(r, db)
	routes.RegisterPatientRoutes(r)
	routes.SetupMessageRoutes(r, db)
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
