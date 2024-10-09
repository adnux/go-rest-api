package models

import "github.com/adnux/go-rest-api/db"

type Registration struct {
	ID      int64 `json:"id"`
	EventId int64 `json:"event_id" binding:"required"`
	UserId  int64 `json:"user_id" binding:"required"`
	Active  bool  `json:"active"`
}

func GetRegistrationsForEvent(eventId int64, userId int64) ([]Registration, error) {
	query := `
	SELECT reg.id, reg.event_id, reg.user_id, reg.active
	  FROM registrations reg
	  LEFT JOIN events ev
		  ON reg.event_id = ev.id
	 WHERE event_id = ?
	   AND ev.user_id = ?
	`
	rows, err := db.DB.Query(query, eventId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registrations []Registration
	for rows.Next() {
		var registration Registration
		if err := rows.Scan(
			&registration.ID,
			&registration.EventId,
			&registration.UserId,
			&registration.Active,
		); err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}

	return registrations, nil
}
