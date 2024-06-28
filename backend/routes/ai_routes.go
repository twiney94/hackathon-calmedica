package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAIRoutes(router *gin.Engine) {
	router.POST("/generate/:patientId", handlers.HandleAI)
}
