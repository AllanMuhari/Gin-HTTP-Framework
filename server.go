package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/test", func(Context *gin.Context) {
		Context.JSON(200, gin.H{
			"message ": "OK !!",
		})
	})

	server.Run(":8080")
}
