package routes

import (
	"github.com/PaulFWatts/rest_api_golang/middlewares"
	"github.com/gin-gonic/gin"
)

	func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // This can be used to get a specific event by ID

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events",  createEvent)
	authenticated.PUT("/events/:id", updateEvent) // This can be used to update a specific event by ID
	authenticated.DELETE("/events/:id", deleteEvent) // This can be used to delete a specific
	authenticated.POST("/events/:id/register", registerForEvent) // This can be used to register for a specific event by ID
	authenticated.DELETE("/events/:id/register",cancelRegistration) // This can be used to unregister from a specific event by ID

	server.POST("/signup", signup) // This can be used to handle user signup
	server.POST("/login", login) // This can be used to handle user login
}