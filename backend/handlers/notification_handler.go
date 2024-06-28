package handlers

import (
	"backend/config"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateNotification crée une nouvelle notification avec l'ID du patient dans l'URL
func CreateNotification(c *gin.Context) {
	patientID := c.Param("patientId")
	var input struct {
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	notification := models.Notification{
		PatientID: patientID,
		Message:   input.Message,
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"notification": notification})
}

// GetAllNotifications récupère toutes les notifications
func GetAllNotifications(c *gin.Context) {
	var notifications []models.Notification
	if err := config.DB.Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}
