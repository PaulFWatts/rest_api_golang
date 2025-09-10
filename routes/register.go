package routes

import (
	"net/http"
	"strconv"

	"github.com/PaulFWatts/rest_api_golang/models"
	"github.com/gin-gonic/gin"
)

// registerForEvent handles POST /events/:id/register - registers authenticated user for an event
//
// This handler allows authenticated users to register for events they don't own.
// It validates the event exists, then creates a registration record in the database.
//
// HTTP Method: POST
// Endpoint: /events/:id/register
// Authentication: Required (JWT middleware validates Authorization header)
// URL Parameters: id (integer) - The unique event identifier
//
// Response Codes:
//   - 201 Created: Successfully registered for event
//   - 400 Bad Request: Invalid event ID format
//   - 500 Internal Server Error: Event not found or database operation failed
//
// Response Body:
//
//	Success: {"message": "Registered!"}
//	Error: {"message": "error description"}
//
// Security Notes:
//   - User ID extracted from JWT token via middleware
//   - Validates event exists before allowing registration
//   - Creates many-to-many relationship in registrations table
func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

// cancelRegistration handles DELETE /events/:id/register - cancels user's event registration
//
// This handler allows authenticated users to cancel their registration for events.
// It removes the registration record from the database without requiring event ownership.
//
// HTTP Method: DELETE
// Endpoint: /events/:id/register
// Authentication: Required (JWT middleware validates Authorization header)
// URL Parameters: id (integer) - The unique event identifier
//
// Response Codes:
//   - 200 OK: Successfully cancelled registration
//   - 400 Bad Request: Invalid event ID format
//   - 500 Internal Server Error: Database operation failed
//
// Response Body:
//
//	Success: {"message": "Cancelled!"}
//	Error: {"message": "error description"}
//
// Security Notes:
//   - User ID extracted from JWT token via middleware
//   - Only removes registration for the authenticated user
//   - Does not require event ownership validation
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
