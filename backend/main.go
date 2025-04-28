package main

import (
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"status":  true,
		})
	})
	router.Run(":8080")

	config := config.LoadConfig()
	log.Println(config)

}
