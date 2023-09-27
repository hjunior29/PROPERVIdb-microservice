package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunior29/PROPERVIdb-microservice/handlers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/", handlers.Home)
	r.POST("/add-properties", handlers.AddProperties)
	r.Run()
}
