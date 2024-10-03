package main

import (
	"net/http"

	"github.com/adnux/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event.ID = 1
	event.UserId = 1

	newEvent, err := event.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"error": "Could not create event",
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   newEvent,
	})

}
