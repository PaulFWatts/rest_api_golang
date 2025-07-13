package routes
import (
		"github.com/gin-gonic/gin"
)

	func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // This can be used to get a specific event by ID
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent) // This can be used to update a specific event by ID
	//server.DELETE("/events/:id", deleteEvent) // This can be used to delete a specific
}