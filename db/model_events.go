package db

import (
	"errors"
)

func (event Event) SaveEvent() (Event, error) {
	event, err := DBQueries.InsertEvent(CTX, InsertEventParams{
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    event.DateTime,
		UserID:      event.UserID,
	})
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func GetAllEvents() ([]Event, error) {
	events, err := DBQueries.GetEvents(CTX)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	event, err := DBQueries.GetEvent(CTX, id)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event *Event) UpdateEvent() error {
	err := DBQueries.UpdateEvent(CTX, UpdateEventParams{
		ID:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    event.DateTime,
		UserID:      event.UserID,
	})
	if err != nil {
		return errors.New("event not found")
	}

	return nil
}

func (event Event) DeleteEvent() error {
	err := DBQueries.DeleteEvent(CTX, event.ID)
	if err != nil {
		return errors.New("event not found")
	}

	return nil
}

func (event Event) Register(userId int64) error {
	registration, err := DBQueries.RegisterUserForEvent(CTX, RegisterUserForEventParams{
		EventID: event.ID,
		UserID:  userId,
	})
	if err != nil {
		return err
	}

	if (registration == Registration{} || registration.ID == 0) {
		return errors.New("already registered")
	}

	return nil
}

func (event Event) CancelRegistration(userId int64) error {
	registration, err := DBQueries.CancelRegistration(CTX, CancelRegistrationParams{
		EventID: event.ID,
		UserID:  userId,
	})
	if err != nil {
		return err
	}

	if (registration == Registration{} || registration.ID == 0) {
		return errors.New("no active registration found")
	}

	return nil
}
