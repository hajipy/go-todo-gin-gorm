package main

import (
	"log"
	"net/http"
)

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
