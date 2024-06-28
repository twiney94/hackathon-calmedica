package handlers

import (
	"backend/config"
	"backend/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPatientByID récupère un patient par son UUID
func GetPatientByID(c *gin.Context) {
	var patient models.Patient
	uuid := c.Param("uuid")

	if err := config.DB.Where("uuid = ?", uuid).First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var patientDTO models.PatientDTO
	patientDTO.UUID = patient.UUID
	patientDTO.Status = patient.Status
	patientDTO.Phone = patient.Phone
	patientDTO.Keywords = patient.Keywords

	c.JSON(http.StatusOK, gin.H{"patient": patientDTO})
}

// GetAllPatients récupère tous les patients et ne renvoie que les champs UUID, Status et Phone
func GetAllPatients(c *gin.Context) {
	var patients []models.Patient
	if err := config.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patients"})
		return
	}

	// Transformer les patients en DTOs
	var patientDTOs []models.PatientDTO
	for _, patient := range patients {
		patientDTOs = append(patientDTOs, models.PatientDTO{
			UUID:     patient.UUID,
			Status:   patient.Status,
			Phone:    patient.Phone,
			Keywords: patient.Keywords,
		})
	}

	c.JSON(http.StatusOK, gin.H{"patients": patientDTOs})
}

func UpdatePatient(c *gin.Context) {
	uuid := c.Param("uuid")

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var patient models.Patient
	if err := config.DB.First(&patient, "uuid = ?", uuid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Handle Keywords field
	if keywords, ok := updateData["keywords"].([]interface{}); ok {
		// Convert the slice of interfaces to a slice of strings
		var keywordsStrings []string
		for _, keyword := range keywords {
			if str, ok := keyword.(string); ok {
				keywordsStrings = append(keywordsStrings, str)
			}
		}
		// Marshal the slice of strings to JSON
		keywordsJSON, err := json.Marshal(keywordsStrings)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid keywords format"})
			return
		}
		updateData["keywords"] = string(keywordsJSON)
	}

	// Update only the provided fields
	if err := config.DB.Model(&patient).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch the updated patient
	if err := config.DB.First(&patient, "uuid = ?", uuid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated patient"})
		return

	}

	var patientDTO models.PatientDTO
	patientDTO.UUID = patient.UUID
	patientDTO.Status = patient.Status
	patientDTO.Phone = patient.Phone
	patientDTO.Keywords = patient.Keywords

	c.JSON(http.StatusOK, gin.H{"patient": patientDTO})
}
