package models

import (
	"errors"

	"github.com/adnux/go-rest-api/db"
	"github.com/adnux/go-rest-api/utils"
)

type User struct {
	ID        int64  `json:"id" gorm:"primarykey"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password,omitempty" binding:"required"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
}

func GetUserByID(id int64) (*User, error) {
	var user User
	result := db.DB.
		Model(&User{}).
		Where("id = ?", id).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (user *User) Save() error {
	hashedPassword, err := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	if err != nil {
		return err
	}

	result := db.DB.
		Model(&User{}).
		Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (user User) DeleteUser() error {
	result := db.DB.
		Model(&User{}).
		Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (user *User) ValidateCredentials() error {
	plainPassword := user.Password
	result := db.DB.
		Model(&User{}).
		Where("email = ?", user.Email).
		First(&user)

	if result.Error != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(plainPassword, user.Password)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.DB.
		Model(User{}).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
