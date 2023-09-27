package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hjunior29/PROPERVIdb-service/db"
	"github.com/hjunior29/PROPERVIdb-service/models"
)

func Home(c *gin.Context) {
	message := "Microservice intended for database operations."
	currentTime := time.Now().Format(time.RFC3339)
	status := "running"

	c.JSON(200, gin.H{
		"message": message,
		"status":  status,
		"time":    currentTime,
	})
}

func AddProperties(c *gin.Context) {
	var properties []models.Properties

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}

	if err := json.Unmarshal(body, &properties); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode JSON payload"})
		return
	}

	if err := db.DB.Create(&properties).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to add properties to the database"})
		log.Printf("Failed to add properties to the database. Payload: %v\n", properties)
		log.Printf("Error: %v\n", err)
		return
	}

	c.JSON(200, gin.H{"message": "Properties added successfully"})
}
