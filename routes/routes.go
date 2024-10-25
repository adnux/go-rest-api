package routes

import (
	"github.com/adnux/go-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	// Events
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	// Registrations
	authenticated.GET("/events/:id/registrations", getAllRegistrationsFromEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.PUT("/events/:id/unregister", cancelRegistration)
	// Users
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.DELETE("user/:id", deleteUser)
}
