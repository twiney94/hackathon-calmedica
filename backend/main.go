package main

import (
	"backend/config"
	"backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDb()
	config.SyncDatabase()
	config.LoadEnvVariables()
}

func main() {
	//db connection result := config.DB.Create(&user)
	db := config.DB

	r := gin.Default()

	// Register routes
	routes.SetupUserRoutes(r, db)

	// Start server
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
