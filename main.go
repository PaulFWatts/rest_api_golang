package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default() // Engine instance with default middleware (logger and recovery)

	server.GET("/events", getEvents)

	server.Run(":8080") // Start the server on port 8080
}

func getEvents(context *gin.Context) {
	// Handler logic for retrieving events
	// This function will be called when a GET request is made to /events
	context.JSON(http.StatusOK, gin.H{
		"message": "List of events",
	})
}
