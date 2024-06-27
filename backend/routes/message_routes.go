package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupMessageRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/messages/:patientId", handlers.GetMessageByPatientId(db))
	r.POST("/messages", handlers.CreateMessage(db))
	r.GET("/messages", handlers.GetMessages(db))
}
