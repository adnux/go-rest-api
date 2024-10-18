package models

import (
	"github.com/adnux/go-rest-api/db"
	"gorm.io/gorm"
)

type Registration struct {
	ID      int64 `json:"id" gorm:"primarykey"`
	EventId int64 `json:"event_id" binding:"required"`
	Active  bool  `json:"active"`
	UserID  int64 `json:"user_id" binding:"required"`
	User    User  `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

func GetRegistrationsForEvent(eventId int64, userId int64) ([]Registration, error) {
	var registrations []Registration
	result := db.DB.
		Model(&Registration{}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, first_name, last_name")
		}).
		Where("event_id = ? AND user_id = ?", eventId, userId).
		Find(&registrations)

	if result.Error != nil {
		return nil, result.Error
	}

	return registrations, nil
}
