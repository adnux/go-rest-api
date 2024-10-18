package models

import (
	"errors"
	"time"

	"github.com/adnux/go-rest-api/db"
	"gorm.io/gorm"
)

type Event struct {
	ID          int64     `json:"id" gorm:"primarykey"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int64     `json:"user_id" binding:"required"`
	User        User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

func (Event) TableName() string {
	return "events"
}

func (event Event) SaveEvent() (Event, error) {

	result := db.DB.
		Model(&Event{}).
		Create(&event)

	if result.Error != nil {
		return Event{}, result.Error
	}

	return event, nil
}

func GetAllEvents() ([]Event, error) {
	var events []Event
	result := db.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, first_name, last_name")
		}).
		Find(&events)

	if result.Error != nil {
		return nil, result.Error
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	var event Event
	result := db.DB.
		Model(&Event{}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, first_name, last_name")
		}).
		Where("id = ?", id).
		First(&event)

	if result.Error != nil {
		return nil, result.Error
	}

	return &event, nil
}

func (event *Event) UpdateEvent() error {
	result := db.DB.
		Model(&Event{ID: event.ID}).
		Where("id = ?", event.ID).
		Updates(&event)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no event found")
	}

	return nil
}

func (event Event) DeleteEvent() error {
	result := db.DB.
		Model(&Event{ID: event.ID}).
		Delete(&event)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no event found")
	}

	return nil
}

func (event Event) Register(userId int64) error {
	user, err := GetUserByID(userId)
	if err != nil {
		return err
	}

	result := db.DB.
		Model(&Registration{}).
		Create(&Registration{EventId: event.ID, User: *user})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (event Event) CancelRegistration(userId int64) error {
	result := db.DB.
		Model(&Registration{}).
		Where("event_id = ? AND user_id = ?", event.ID, userId).
		Update("active", false)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no active registration found")
	}

	return nil
}

func (event *Event) SetUser(authUserId int64) (*Event, error) {
	user, err := GetUserByID(authUserId)
	if err != nil {
		return &Event{}, err
	}

	event.UserID = user.ID
	event.User = *user

	return event, nil
}

func (event Event) ToJSON() string {
	return `{
		"id": ` + string(event.ID) + `,
		"name": "` + event.Name + `",
		"description": "` + event.Description + `",
		"location": "` + event.Location + `",
		"datetime": "` + event.DateTime.String() + `"
		"user_id": ` + string(event.UserID) + `
	}`
}
