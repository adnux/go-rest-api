package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/adnux/go-rest-api/db"
	"github.com/adnux/go-rest-api/utils"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (user User) Save() (User, error) {
	fmt.Println("User.Save() called", user)
	hashedPassword, err := utils.HashPassword(user.Password)
	fmt.Println("hashedPassword", hashedPassword)

	if err != nil {
		return User{}, err
	}

	// Check if db.DBQueries is nil
	if db.DBQueries == nil {
		return User{}, errors.New("db.DBQueries is nil")
	}
	fmt.Println("DBQueries", db.DBQueries)

	// Check if db.CTX is nil
	if db.CTX == nil {
		return User{}, errors.New("db.CTX is nil")
	}
	fmt.Println("DB.CTX", db.CTX)

	insertedUser, err := db.DBQueries.InsertUser(db.CTX, db.InsertUserParams{
		Email:     user.Email,
		Password:  hashedPassword,
		FirstName: sql.NullString{String: user.FirstName, Valid: user.FirstName != ""},
		LastName:  sql.NullString{String: user.LastName, Valid: user.LastName != ""},
	})

	user.ID = insertedUser.ID
	user.Email = insertedUser.Email
	user.Password = insertedUser.Password
	user.FirstName = insertedUser.FirstName.String
	user.LastName = insertedUser.LastName.String
	return user, err
}

func (user User) DeleteUser() error {
	query := `
	DELETE FROM users WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.ID)

	return err
}

func (user *User) ValidateCredentials() error {
	userFound, err := db.DBQueries.GetUserByEmail(db.CTX, user.Email)

	if err != nil {
		return errors.New("credentials invalid")
	}
	retrievedPassword := userFound.Password

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
