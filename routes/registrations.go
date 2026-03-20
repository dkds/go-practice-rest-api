package routes

import (
	"net/http"
	"strconv"

	"dkds.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
			"error":   err.Error(),
		})
		return
	}

	err = event.RegisterUser(context.GetInt64("userId"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register to event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Registered to event successfully",
		"event":   event,
	})
}

func cancelRegistration(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
			"error":   err.Error(),
		})
		return
	}

	err = event.CancelRegistration(context.GetInt64("userId"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel registration to event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Registration to event canceled successfully",
		"event":   event,
	})
}
