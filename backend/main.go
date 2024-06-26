package main

import (
	"backend/config"
	"backend/handlers"
	"backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDb()
	config.SyncDatabase()
	config.LoadEnvVariables()
}

func main() {
	db := config.DB

	r := gin.Default()

	// Register user routes
	routes.SetupUserRoutes(r, db)

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
