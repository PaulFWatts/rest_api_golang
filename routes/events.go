package routes

import (
	"net/http"
	"strconv"

	"github.com/PaulFWatts/rest_api_golang/models"
	"github.com/gin-gonic/gin"
)

// getEvents handles GET /events - retrieves all events from the database
//
// This handler fetches all events using the models.GetAllEvents() function
// and returns them as a JSON array. If database errors occur, it returns
// a 500 Internal Server Error with a generic error message.
//
// HTTP Method: GET
// Endpoint: /events
// Authentication: Not required
//
// Response Codes:
//   - 200 OK: Successfully retrieved events (returns JSON array)
//   - 500 Internal Server Error: Database query failed
//
// Response Body:
//
//	Success: Array of Event objects
//	Error: {"message": "Could not fetch events. Try again later."}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

// getEvent handles GET /events/:id - retrieves a specific event by ID
//
// This handler parses the event ID from the URL parameter, validates it,
// and fetches the corresponding event from the database. The ID must be
// a valid integer that can be parsed as int64.
//
// HTTP Method: GET
// Endpoint: /events/:id
// Authentication: Not required
// URL Parameters: id (integer) - The unique event identifier
//
// Response Codes:
//   - 200 OK: Successfully retrieved the event
//   - 400 Bad Request: Invalid or non-numeric event ID
//   - 500 Internal Server Error: Database query failed or event not found
//
// Response Body:
//
//	Success: Single Event object
//	Error: {"message": "Could not parse event id."} or {"message": "Could not fetch event."}
func getEvent(context *gin.Context) {
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

	context.JSON(http.StatusOK, event)
}

// createEvent handles POST /events - creates a new event with JWT authentication
//
// This handler creates a new event using the authenticated user's ID from the
// JWT middleware. The middleware validates the token and sets the user ID in
// the Gin context, which is then used for proper event ownership.
//
// HTTP Method: POST
// Endpoint: /events
// Authentication: Required (JWT middleware validates Authorization header)
//
// Request Headers:
//
//	Authorization: JWT token string (validated by middleware)
//
// Request Body: JSON object with required fields:
//   - name (string): Event name
//   - description (string): Event description
//   - location (string): Event location
//   - dateTime (string): Event date/time in ISO format
//
// Response Codes:
//   - 201 Created: Event successfully created
//   - 400 Bad Request: Invalid JSON request data
//   - 401 Unauthorized: Missing, invalid, or expired JWT token (handled by middleware)
//   - 500 Internal Server Error: Database save operation failed
//
// Response Body:
//
//	Success: {"message": "Event created!", "event": {...}}
//	Error: {"message": "error description"}
//
// Security Notes:
//   - Authentication handled by JWT middleware before reaching this handler
//   - User ID automatically extracted from validated JWT token via context
//   - Uses JSON binding with struct validation tags
//   - Proper event ownership through authenticated user ID
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userID := context.GetInt64("userId")
	event.UserID = userID

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

// updateEvent handles PUT /events/:id - updates an existing event
//
// This handler validates the event ID from URL parameters, verifies the event
// exists in the database, parses the JSON update data, and saves the changes.
// The event ID from the URL takes precedence over any ID in the request body.
//
// HTTP Method: PUT
// Endpoint: /events/:id
// Authentication: Not currently required (TODO: add JWT validation)
// URL Parameters: id (integer) - The unique event identifier to update
//
// Request Body: JSON object with fields to update:
//   - name (string): Updated event name
//   - description (string): Updated event description
//   - location (string): Updated event location
//   - dateTime (string): Updated event date/time in ISO format
//
// Response Codes:
//   - 200 OK: Event successfully updated
//   - 400 Bad Request: Invalid event ID or JSON request data
//   - 500 Internal Server Error: Event not found or database operation failed
//
// Response Body:
//
//	Success: {"message": "Event updated successfully!"}
//	Error: {"message": "error description"}
//
// Security Notes:
//   - Validates event exists before allowing updates
//   - URL parameter ID overrides request body ID
//   - TODO: Add authentication to prevent unauthorized updates
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

// deleteEvent handles DELETE /events/:id - removes an event from the database
//
// This handler validates the event ID from URL parameters, verifies the event
// exists in the database, and then permanently deletes it. This operation
// cannot be undone once completed.
//
// HTTP Method: DELETE
// Endpoint: /events/:id
// Authentication: Not currently required (TODO: add JWT validation)
// URL Parameters: id (integer) - The unique event identifier to delete
//
// Response Codes:
//   - 200 OK: Event successfully deleted
//   - 400 Bad Request: Invalid or non-numeric event ID
//   - 500 Internal Server Error: Event not found or database deletion failed
//
// Response Body:
//
//	Success: {"message": "Event deleted successfully!"}
//	Error: {"message": "error description"}
//
// Security Notes:
//   - Validates event exists before attempting deletion
//   - Permanent operation - no soft delete implemented
//   - TODO: Add authentication and authorization checks
//   - TODO: Consider adding ownership validation (user can only delete own events)
//
// Database Impact:
//   - Removes record permanently from events table
//   - Foreign key constraints should be considered for related data
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
