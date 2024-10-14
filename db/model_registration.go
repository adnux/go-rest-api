package db

// type Registration struct {
// 	ID      int64 `json:"id"`
// 	EventId int64 `json:"event_id" binding:"required"`
// 	UserId  int64 `json:"user_id" binding:"required"`
// 	Active  bool  `json:"active"`
// }

func GetRegistrationsForEvent(eventId int64, userId int64) ([]GetRegistrationsRow, error) {
	registrations, err := DBQueries.GetRegistrations(CTX, GetRegistrationsParams{
		EventID: eventId,
		UserID:  userId,
	})
	if err != nil {
		return nil, err
	}
	return registrations, nil
}
