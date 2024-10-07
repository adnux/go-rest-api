package routes

import (
	"net/http"
	"strconv"

	"github.com/adnux/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not fetch events. Try again later.",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse event id.",
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not fetch event.",
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
			"error":   err.Error(),
			"message": "Could not bind event data.",
		})
		return
	}

	event, err = event.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Could not create event.",
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse event id.",
		})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Event not found.",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not bind event data.",
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not update event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated!",
		"event":   updatedEvent,
	})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse event id.",
		})
		return
	}

	eventToDelete, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Event not found.",
		})
		return
	}

	err = eventToDelete.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not delete event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted!",
	})
}

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not bind user data.",
		})
		return
	}

	user, err = user.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Could not create user.",
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created!",
		"user":    user,
	})
}