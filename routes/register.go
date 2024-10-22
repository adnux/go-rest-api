package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adnux/go-rest-api/models"
	"github.com/adnux/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func registerForEvent(w http.ResponseWriter, r *http.Request) {
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
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not fetch event.",
		)
		return
	}

	err = event.Register(authUserId)

	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not register user for event.",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	jsonData, err := json.Marshal(gin.H{
		"message": "Registered to the event!",
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func cancelRegistration(w http.ResponseWriter, r *http.Request) {
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

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(authUserId)

	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not cancel registration.",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(gin.H{
		"message": "Registration cancelled!",
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getAllRegistrationsFromEvent(w http.ResponseWriter, r *http.Request) {
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

	registrations, err := models.GetRegistrationsForEvent(eventId, authUserId)

	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not fetch registrations.",
		)
		return
	}

	jsonData, err := json.Marshal(registrations)

	if err != nil {
		utils.CreateHttpErrorMessage(
			w,
			http.StatusInternalServerError,
			err,
			"Could not marshal registrations data.",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
