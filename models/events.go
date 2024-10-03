package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (event Event) Save() ([]Event, error) {

	event.ID = len(events) + 1
	event.UserId = 1

	// add to database
	events = append(events, event)
	return events, nil
}

func GetAllEvents() ([]Event, error) {
	// get from database
	return events, nil
}
