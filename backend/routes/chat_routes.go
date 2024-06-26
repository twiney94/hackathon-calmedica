package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupChatRoutes(r *gin.Engine) {
	r.GET("/ws", handlers.HandleConnections)
}
