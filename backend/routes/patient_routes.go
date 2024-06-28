package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterPatientRoutes(router *gin.Engine) {
	router.GET("/patients/:uuid", handlers.GetPatientByID) // Utilisez cette route pour récupérer un patient par UUID
	router.GET("/patients", handlers.GetAllPatients)       // Utilisez cette route pour récupérer tous les patients
	router.PUT("/patients/:uuid", handlers.UpdatePatient)  // Utilisez cette route pour mettre à jour un patient
}
