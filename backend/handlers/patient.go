package handlers

import (
	"backend/config"
	"backend/models"
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
			UUID:   patient.UUID,
			Status: patient.Status,
			Phone:  patient.Phone,
			//CreatedAt: patient.CreatedAt,
			//UpdatedAt: patient.UpdatedAt,
			//DeletedAt: patient.DeletedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"patients": patientDTOs})
}
