package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupMessageRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/messages", handlers.CreateMessage)
	r.GET("/messages", handlers.GetMessages)
}
