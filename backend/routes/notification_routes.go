package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(router *gin.Engine) {
	router.POST("/patients/:patientId/notifications", handlers.CreateNotification)
	router.GET("/notifications", handlers.GetAllNotifications)
}
