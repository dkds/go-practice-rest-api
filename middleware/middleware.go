package middleware

import (
	"net/http"

	"dkds.com/rest-api/security"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"error":   "Invalid credentials",
		})
		return
	}

	err := security.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	userId, err := security.ExtractUserIdFromToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
