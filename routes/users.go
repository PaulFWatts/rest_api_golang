package routes

import (
	"net/http"

	"github.com/PaulFWatts/rest_api_golang/models"
	"github.com/gin-gonic/gin"
)

// signup validates the input, creates a new user, and saves it to the database
func signup(context *gin.Context) {
	
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user_id": user.ID})
}
