package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAI(c *gin.Context) {
	// Récupérer le fichier du formulaire
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Erreur lors de la récupération du fichier : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Impossible de récupérer le fichier"})
		return
	}
	defer file.Close()

	// Préparer le buffer pour le form-data
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Ajouter le fichier au form-data
	part, err := writer.CreateFormFile("file", header.Filename)
	if err != nil {
		log.Printf("Erreur lors de la création du form-data : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du form-data"})
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Printf("Erreur lors de l'ajout du fichier au form-data : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'ajout du fichier au form-data"})
		return
	}
	writer.Close()

	// Créer et envoyer la requête POST
	req, err := http.NewRequest("POST", "http://localhost:5000/whisper", &buffer)
	if err != nil {
		log.Printf("Erreur lors de la création de la requête : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de la requête"})
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Erreur lors de l'envoi de la requête : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'envoi de la requête"})
		return
	}
	defer resp.Body.Close()

	// Lire la réponse du serveur
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erreur lors de la lecture de la réponse : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture de la réponse"})
		return
	}

	// Décoder la réponse JSON pour gérer les caractères échappés
	var decodedResponse map[string]interface{}
	err = json.Unmarshal(body, &decodedResponse)
	if err != nil {
		log.Printf("Erreur lors du décodage de la réponse JSON : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du décodage de la réponse JSON"})
		return
	}

	// Réponse du serveur
	if resp.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"message": "Fichier traité avec succès", "response": decodedResponse})
	} else {
		log.Printf("Erreur lors du traitement du fichier, status code: %d, réponse: %s", resp.StatusCode, string(body))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du traitement du fichier", "status_code": resp.StatusCode, "response": decodedResponse})
	}
}
