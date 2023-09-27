package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunior29/PROPERVIdb-service/handlers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/", handlers.Home)
	r.POST("/add-properties", handlers.AddProperties)
	r.Run()
}
