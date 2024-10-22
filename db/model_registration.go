package db

import "errors"

func GetRegistrationsForEvent(eventId int64, userId int64) ([]GetRegistrationsRow, error) {
	registrations, err := DBQueries.GetRegistrations(CTX, GetRegistrationsParams{
		EventID: eventId,
		UserID:  userId,
	})
	if err != nil {
		return nil, errors.New("registrations not found")
	}
	return registrations, nil
}
