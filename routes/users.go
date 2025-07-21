package routes

import (
	"net/http"

	"github.com/PaulFWatts/rest_api_golang/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	// Handle user signup logic here
	// This function should validate the input, create a new user, and save it to the database
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
