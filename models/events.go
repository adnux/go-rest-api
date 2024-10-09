package models

import (
	"time"

	"github.com/adnux/go-rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserId      int64     `json:"user_id"  binding:"required"`
}

func (event Event) SaveEvent() (Event, error) {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return Event{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId,
	)
	if err != nil {
		return Event{}, err
	}

	event.ID, err = result.LastInsertId()

	return event, err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId,
		event.ID,
	)

	return err
}

func (event Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
