package models

import (
	"errors"

	"github.com/PaulFWatts/rest_api_golang/db"
	"github.com/PaulFWatts/rest_api_golang/utils"
)

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

// ValidateCredentials verifies a user's email and password against the database
//
// This method performs the following steps:
// 1. Queries the database for a user with the provided email address
// 2. Retrieves the stored hashed password for that user
// 3. Compares the provided plain text password with the stored hash
//
// Parameters:
// - Uses the User struct's Email and Password fields for validation
//
// Returns:
// - nil if the credentials are valid (user exists and password matches)
// - error with message "Credentials invalid" if:
//   - No user found with the provided email address
//   - The provided password doesn't match the stored password hash
//
// Note: This method updates the User's ID field with the database ID if validation succeeds
func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
