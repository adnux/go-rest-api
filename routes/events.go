package routes

import (
	"net/http"
	"strconv"

	"github.com/adnux/go-rest-api/db"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := db.GetAllEvents()
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

	event, err := db.GetEventByID(eventId)

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
	var event db.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not bind event data.",
		})
		return
	}

	event, err = event.SaveEvent()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Could not create event.",
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
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

	authUserId := context.GetInt64("authUserId")
	event, err := db.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not fetch the event.",
		})
		return
	}

	if event.UserID != authUserId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update event.",
		})
		return
	}

	var updatedEvent db.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse request data.",
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

	authUserId := context.GetInt64("authUserId")
	event, err := db.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not fetch the event.",
		})
		return
	}

	if event.UserID != authUserId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to delete event.",
		})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not delete the event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully!",
	})
}
