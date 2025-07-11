package main

import (
	"github.com/PaulFWatts/rest_api_golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default() // Engine instance with default middleware (logger and recovery)

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // Start the server on port 8080
}

func getEvents(context *gin.Context) {
	// Handler logic for retrieving events
	// This function will be called when a GET request is made to /events
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// Handler logic for creating a new event
	// This function will be called when a POST request is made to /events
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
