package main

import (
	"net/http"
	"strconv"

	"github.com/PaulFWatts/rest_api_golang/db"
	"github.com/PaulFWatts/rest_api_golang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database connection and create necessary tables
	server := gin.Default() // Engine instance with default middleware (logger and recovery)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // This can be used to get a specific event by ID
	server.POST("/events", createEvent)

	server.Run(":8080") // Start the server on port 8080
}

func getEvents(context *gin.Context) {
	// Handler logic for retrieving events
	// This function will be called when a GET request is made to /events
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	// Handler logic for retrieving a specific event by ID
	// This function will be called when a GET request is made to /events/:id
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}
	event, err := models.GetEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve event."})
		return
	}
	context.JSON(http.StatusOK, event) // Return the event as JSON response
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

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
