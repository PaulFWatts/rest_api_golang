package models

import "github.com/PaulFWatts/rest_api_golang/db"
import "github.com/PaulFWatts/rest_api_golang/utils"

// User represents a user in the system
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save inserts a new user into the database
func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}
