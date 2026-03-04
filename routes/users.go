package routes

import (
	"net/http"

	"dkds.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func loginUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request",
			"error":   err.Error(),
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login success",
	})
}

func signupUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request",
			"error":   err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"event":   user,
	})
}
