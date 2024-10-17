package db

import (
	"errors"

	"github.com/adnux/go-rest-api/utils"
)

func (user User) Save() (User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return User{}, errors.New("could not hash password")
	}

	insertedUser, err := DBQueries.InsertUser(CTX, InsertUserParams{
		Email:     user.Email,
		Password:  hashedPassword,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})

	user.ID = insertedUser.ID
	user.Email = insertedUser.Email
	user.Password = insertedUser.Password
	user.FirstName = insertedUser.FirstName
	user.LastName = insertedUser.LastName
	return user, err
}

func (user User) DeleteUser() error {
	err := DBQueries.DeleteUser(CTX, user.ID)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (user *User) ValidateCredentials() error {
	userFound, err := DBQueries.GetUserByEmail(CTX, user.Email)

	if err != nil {
		return errors.New("credentials invalid")
	}
	retrievedPassword := userFound.Password

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	user.ID = userFound.ID

	return nil
}
