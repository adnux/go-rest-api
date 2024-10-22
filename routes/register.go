package routes

import (
	"net/http"
	"strconv"

	"github.com/adnux/go-rest-api/db"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	authUserId := context.GetInt64("authUserId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse event id.",
		})
		return
	}

	event, err := db.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not fetch event.",
		})
		return
	}

	err = event.Register(authUserId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not register user for event.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Registered to the event!",
	})
}

func cancelRegistration(context *gin.Context) {
	authUserId := context.GetInt64("authUserId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse event id.",
		})
		return
	}

	var event db.Event
	event.ID = eventId

	err = event.CancelRegistration(authUserId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not cancel registration.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Registration cancelled!",
	})
}

func getAllRegistrationsFromEvent(context *gin.Context) {
	authUserId := context.GetInt64("authUserId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse event id.",
		})
		return
	}

	registrations, err := db.GetRegistrationsForEvent(eventId, authUserId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not fetch registrations.",
		})
		return
	}

	context.JSON(http.StatusOK, registrations)
}
