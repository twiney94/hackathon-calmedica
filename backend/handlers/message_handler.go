package handlers

import (
	"backend/config"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateMessage handles the creation of a new message
func CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message"})
		return
	}

	c.JSON(http.StatusOK, message)
}

// GetMessages handles retrieving messages
func GetMessages(c *gin.Context) {
	var messages []models.Message
	if err := config.DB.Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
