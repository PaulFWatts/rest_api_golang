package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword returns the hashed password and any error encountered
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err 
}