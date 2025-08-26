package main

import (


	"github.com/PaulFWatts/rest_api_golang/db"
	"github.com/PaulFWatts/rest_api_golang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database connection and create necessary tables
	server := gin.Default() // Engine instance with default middleware (logger and recovery)

	routes.RegisterRoutes(server) // Register the routes defined in the routes package

	server.Run(":8080") // Start the server on port 8080
}


