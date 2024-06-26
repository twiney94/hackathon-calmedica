package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	user := router.Group("/user")
	{
		user.GET("/", handlers.GetUser())
	}
}
