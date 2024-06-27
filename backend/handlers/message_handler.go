package handlers

import (
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageInput struct {
	PatientId string `json:"patient_id"`
	Content   string `json:"content"`
}

func CreateMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input MessageInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		message := models.Message{
			PatientId: input.PatientId,
			Content:   input.Content,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&message).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": message})
	}
}

func GetMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var messages []models.Message
		if err := db.Find(&messages).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": messages})
	}
}

func GetMessageByPatientId(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		patientId := c.Param("patientId")

		var messages []models.Message
		var patient models.Patient

		if err := db.Where("UUID = ?", patientId).First(&patient).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		if err := db.Where("patient_id = ?", patient.UUID).Find(&messages).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": messages})
	}
}
