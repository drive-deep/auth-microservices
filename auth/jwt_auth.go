package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
// HashPasswordWithFixedSalt hashes the password with a fixed salt and returns the hashed password
func HashPasswordWithSalt(password, salt string) (string, error) {
	// Trim any leading or trailing whitespace from both the password and fixed salt
	password = strings.TrimSpace(password)
	salt = strings.TrimSpace(salt)

	// Combine the password and salt
	passwordWithSalt := password + salt

	// Create a new SHA-256 hash
	hash := sha256.New()
	_, err := hash.Write([]byte(passwordWithSalt))
	if err != nil {
		log.Printf("Error creating SHA-256 hash: %v", err)
		return "", err
	}

	// Get the final hash as a hexadecimal string
	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	// Return the hashed password
	return hashedPassword, nil
}
