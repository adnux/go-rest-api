package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adnux/go-rest-api/models"
	"github.com/adnux/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(w http.ResponseWriter, r *http.Request) {
	events, err := models.GetAllEvents()
	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not fetch events. Try again later.",
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(events)
	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not marshal events data.",
		)
		return
	}
	w.Write(jsonData)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		http.Error(w, "Could not fetch event.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(event)
	if err != nil {
		http.Error(w, "Could not marshal event data.", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusBadRequest,
			err,
			"Could not bind event data.",
		)
		return
	}

	event, err = event.SaveEvent()

	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not create event.",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	jsonData, err := json.Marshal(event)
	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not marshal event data.",
		)
		return
	}
	w.Write(jsonData)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	authUserId, err := strconv.ParseInt(r.Header.Get("authUserId"), 10, 64)

	if err != nil {
		http.Error(w, "Invalid auth user ID", http.StatusBadRequest)
		return
	}

	eventId, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		http.Error(w, "Could not fetch the event.", http.StatusInternalServerError)
		return
	}

	if event.UserId != authUserId {
		http.Error(w, "Not authorized to update event.", http.StatusUnauthorized)
		return
	}

	var updatedEvent models.Event
	err = json.NewDecoder(r.Body).Decode(&updatedEvent)

	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusBadRequest,
			err,
			"Could not parse request data.",
		)
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		utils.CreateHttpErrorMessage(w, http.StatusInternalServerError, err, "Could not update event.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(updatedEvent)
	if err != nil {
		utils.CreateHttpErrorMessage(w, http.StatusInternalServerError, err, "Could not marshal event data.")
		return
	}
	w.Write(jsonData)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	authUserId, err := strconv.ParseInt(r.Header.Get("authUserId"), 10, 64)

	if err != nil {
		http.Error(w, "Invalid auth user ID", http.StatusBadRequest)
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not fetch the event.",
		)
		return
	}

	if event.UserId != authUserId {
		http.Error(w, "Not authorized to delete event.", http.StatusUnauthorized)
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		http.Error(w, "Could not delete the event.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(gin.H{
		"message": "Event deleted successfully!",
	})
	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not marshal response data.",
		)
		return
	}
	w.Write(jsonData)
}
