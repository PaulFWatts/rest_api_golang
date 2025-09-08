package middlewares

import (
	"net/http"

	"github.com/PaulFWatts/rest_api_golang/utils"
	"github.com/gin-gonic/gin"
)

// Authenticate is a Gin middleware that validates JWT tokens and extracts user information
//
// This middleware intercepts HTTP requests to protected endpoints, validates the JWT token
// from the Authorization header, and sets the authenticated user's ID in the Gin context
// for use by downstream handlers.
//
// Security Features:
//   - Validates JWT token signature and expiration
//   - Extracts user ID from token claims
//   - Sets user ID in Gin context for handler access
//   - Aborts request chain if authentication fails
//
// Request Headers Required:
//
//	Authorization: Valid JWT token string
//
// Context Values Set:
//
//	"userId" (int64): The authenticated user's ID extracted from the token
//
// Response Codes:
//   - Continues to next handler if authentication succeeds
//   - 401 Unauthorized: Missing or invalid JWT token
//
// Usage:
//
//	Apply to routes requiring authentication:
//	server.POST("/events", middlewares.Authenticate, createEvent)
//
// Error Handling:
//   - Generic "Not authorized." message to prevent information disclosure
//   - Uses AbortWithStatusJSON to stop the request chain
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
