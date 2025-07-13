package routes
import (
	"net/http"
	"strconv"
	"github.com/PaulFWatts/rest_api_golang/models"
	"github.com/gin-gonic/gin"
)

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

func updateEvent(context *gin.Context) {
	// Handler logic for updating an existing event
	// This function will be called when a PUT request is made to /events/:id
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	_, err = models.GetEvent(eventID)	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve event."})
		return
	}

	var updatedEvent models.Event // Create a new Event instance to hold the updated data
	err = context.ShouldBindJSON(&updatedEvent) // Bind the JSON request body to the updatedEvent struct
		if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventID // Set the ID of the updated event to the one being updated
	
	err = updatedEvent.Update() // Call the Update method to save changes to the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	// Handler logic for deleting an event
	// This function will be called when a DELETE request is made to /events/:id
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

	err = event.Delete() // Call the DeleteEvent method to remove the event from the database
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}