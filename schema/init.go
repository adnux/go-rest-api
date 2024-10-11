package schemas

import (
	"github.com/adnux/go-rest-api/db"
	"github.com/adnux/go-rest-api/models"
)

func CreateTables() {
	createUsersTable()
	createEventsTable()
	createRegistrationsTable()
}

func createUsersTable() {
	err := db.DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("Could not create users table.")
	}
}

func createEventsTable() {
	err := db.DB.AutoMigrate(&models.Event{})

	if err != nil {
		panic("Could not create events table.")
	}
}

func createRegistrationsTable() {
	err := db.DB.AutoMigrate(&models.Registration{})

	if err != nil {
		panic("Could not create registrations table.")
	}
}
