package routes
import (
		"github.com/gin-gonic/gin"
)

	func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // This can be used to get a specific event by ID
	server.POST("/events", createEvent)
}