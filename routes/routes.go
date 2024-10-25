package routes

import (
	"net/http"

	"github.com/adnux/go-rest-api/middlewares"
)

func AuthenticatedRoutes(server *http.ServeMux) {
	// Events
	server.Handle("POST /events", middlewares.EnsureAuthHandler(createEvent))
	server.Handle("PUT /events/{id}", middlewares.EnsureAuthHandler(updateEvent))
	server.Handle("DELETE /events/{id}", middlewares.EnsureAuthHandler(deleteEvent))
	// Registrations
	server.Handle("GET /events/{id}/registrations", middlewares.EnsureAuthHandler(getAllRegistrationsFromEvent))
	server.Handle("POST /events/{id}/register", middlewares.EnsureAuthHandler(registerForEvent))
	server.Handle("PUT /events/{id}/unregister", middlewares.EnsureAuthHandler(cancelRegistration))
}

func RegisterRoutes(server *http.ServeMux) {
	http.Handle("/", server)
	// Events
	server.HandleFunc("GET /events", getEvents)
	server.HandleFunc("GET /events/{id}", getEvent)
	// Users
	server.HandleFunc("POST /signup", signUp)
	server.HandleFunc("POST /login", login)
	server.HandleFunc("DELETE user/{id}", deleteUser)
	// Authenticated routes
	AuthenticatedRoutes(server)
}
