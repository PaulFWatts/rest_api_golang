package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// secretKey is the key used for signing and verifying JWT tokens
// TODO: Move to environment variable for production security
const secretKey = "your_secret_key"

// GenerateToken creates a new JWT token for user authentication
//
// This function generates a JWT token containing the user's email and ID,
// with an expiration time of 2 hours from the current time. The token uses
// HMAC-SHA256 signing method for security.
//
// Parameters:
//   - email: The user's email address to embed in the token claims
//   - userId: The user's unique ID to embed in the token claims
//
// Returns:
//   - string: The signed JWT token string ready for HTTP headers
//   - error: Any error that occurred during token generation or signing
//
// Token Claims:
//   - email: User's email address
//   - userId: User's unique identifier
//   - exp: Token expiration timestamp (2 hours from creation)
//
// Example usage:
//
//	token, err := GenerateToken("user@example.com", 123)
//	if err != nil {
//	    log.Printf("Token generation failed: %v", err)
//	    return
//	}
//	// Use token in Authorization header: "Bearer " + token
//
// Note: Currently the claims are set to empty strings - this should be fixed
// to use the actual email and userId parameters for proper functionality.
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  "",
		"userId": "",
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

// VerifyToken validates and parses a JWT token
//
// This function verifies the signature and validity of a JWT token,
// ensuring it was signed with the correct secret key and uses HMAC signing method.
// It performs comprehensive validation including signature verification, expiration
// checking, and algorithm validation to prevent security vulnerabilities.
//
// Parameters:
//   - token: The JWT token string to verify
//
// Returns:
//   - error: nil if token is valid, otherwise an error describing the validation failure
//     Possible errors include "invalid token" for parsing/signature failures,
//     or "unexpected signing method" for algorithm confusion attacks
//
// Security Features:
//   - Validates HMAC signing method to prevent algorithm confusion attacks
//   - Automatically checks token expiration via jwt.Parse
//   - Verifies token signature against the secret key
//   - Returns generic "invalid token" errors to prevent information disclosure
//
// Example usage:
//
//	err := VerifyToken(tokenString)
//	if err != nil {
//	    // Token is invalid - deny access
//	    return
//	}
//	// Token is valid - proceed with request
func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // Ensure token method is HMAC
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return errors.New("invalid token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	//	return errors.New("invalid token claims")

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}
