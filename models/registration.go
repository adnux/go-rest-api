package models

type Registration struct {
	ID      int64 `json:"id"`
	EventId int64 `json:"event_id" binding:"required"`
	UserId  int64 `json:"user_id" binding:"required"`
	Active  bool  `json:"active"`
}
