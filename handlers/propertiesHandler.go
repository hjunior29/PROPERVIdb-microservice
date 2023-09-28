package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hjunior29/PROPERVIdb-microservice/db"
	"github.com/hjunior29/PROPERVIdb-microservice/models"
)

func Status(c *gin.Context) {
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
	fmt.Println(string(body))
	if err != nil {
		c.JSON(400, gin.H{
			"error":   "Failed to read request body",
			"message": err.Error(),
		})
		return
	}

	if err := json.Unmarshal(body, &properties); err != nil {
		c.JSON(400, gin.H{
			"error":   "Failed to decode JSON payload",
			"message": err.Error(),
		})
		log.Printf("Body: %v\n", body)
		return
	}

	if err := db.DB.Create(&properties).Error; err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to add properties to the database",
			"message": err.Error(),
		})
		log.Printf("Failed to add properties to the database. Payload: %v\n", properties)
		log.Printf("Error: %v\n", err)
		return
	}

	c.JSON(200, gin.H{"message": "Properties added successfully"})

}
