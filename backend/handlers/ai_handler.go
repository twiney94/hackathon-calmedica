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

	// Extraire le transcript
	transcripts, ok := decodedResponse["results"].([]interface{})
	if !ok || len(transcripts) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Aucun transcript trouvé dans la réponse"})
		return
	}

	transcriptData, ok := transcripts[0].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Format du transcript invalide"})
		return
	}

	transcript, ok := transcriptData["transcript"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transcript non trouvé ou format invalide"})
		return
	}

	// Construire le JSON pour la nouvelle requête
	requestBody := map[string]interface{}{
		"model": "mistral",
		"response_format": map[string]string{
			"type": "json_object",
		},
		"system": "Résume le texte que je te donne qui est la transcription d'un appel entre un patient et le personnel médical. En fonction de la gravité, tu vas choisir une couleur entre vert orange et rouge. Je veux également que tu me donnes les mots-clés. Ne me renvoie pas ce que je te donne je veux le résumé. Je veux que tu répondes en français obligatoirement, sous forme d'un JSON avec les 3 clés, COLOR, DESC et KEYWORDS.",
		"prompt": "[INST] " + transcript + " [/INST]",
		"stream": false,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Erreur lors de la création du JSON pour la nouvelle requête : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du JSON pour la nouvelle requête"})
		return
	}

	// Envoyer la nouvelle requête POST
	newReq, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Printf("Erreur lors de la création de la nouvelle requête : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de la nouvelle requête"})
		return
	}
	newReq.Header.Set("Content-Type", "application/json")

	newResp, err := client.Do(newReq)
	if err != nil {
		log.Printf("Erreur lors de l'envoi de la nouvelle requête : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'envoi de la nouvelle requête"})
		return
	}
	defer newResp.Body.Close()

	// Lire la réponse de la nouvelle requête
	newBody, err := io.ReadAll(newResp.Body)
	if err != nil {
		log.Printf("Erreur lors de la lecture de la réponse de la nouvelle requête : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture de la réponse de la nouvelle requête"})
		return
	}

	// Retourner la réponse finale
	c.Data(newResp.StatusCode, "application/json", newBody)
}
