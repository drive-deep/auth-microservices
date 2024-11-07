package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Replace with a secure key

// GenerateToken generates a new JWT token for a given user ID and email.
func GenerateToken(userID, email string, expirationHours int) (string, error) {
	// Create a new token with the specified claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * time.Duration(expirationHours)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// RefreshToken generates a new token by extending the expiration of an existing valid token
func RefreshToken(oldTokenString string, newExpirationHours int) (string, error) {
	// Parse the old token without validating the expiration (since it might be expired)
	token, err := jwt.Parse(oldTokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	// Check if parsing and validation were successful
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract claims and create a new token with updated expiration
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not parse token claims")
	}

	// Set a new expiration time
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(newExpirationHours)).Unix()

	// Create a new token with the updated claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newToken.SignedString(jwtSecret)
}

// generateSalt generates a random 16-byte salt
func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// HashPasswordWithSalt hashes the password with the provided salt and returns the hashed password
func HashPasswordWithSalt(password, salt string) (string, error) {
	// Combine the password and salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}

	return string(hashedPassword), nil
}
