package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunior29/PROPERVIdb-microservice/handlers"
)

func HandleRequest() *gin.Engine {
	r := gin.Default()
	r.GET("/", handlers.Status)
	r.POST("/add-properties", handlers.AddProperties)
	
	return r
}
