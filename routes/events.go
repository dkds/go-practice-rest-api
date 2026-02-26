package routes

import (
	"net/http"
	"strconv"

	"dkds.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
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

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request",
			"error":   err.Error(),
		})
		return
	}

	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
			"error":   err.Error(),
		})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request",
			"error":   err.Error(),
		})
		return
	}

	err = event.Update(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})
}

func deleteEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
			"error":   err.Error(),
		})
		return
	}

	err = models.DeleteEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
