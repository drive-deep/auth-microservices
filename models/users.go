package models

import (
	"time"

	"github.com/drive-deep/auth-microservices/auth"
	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id" pg:"id,pk"`              // Primary key as UUID
	Email     string    `json:"email" pg:"email,unique"`    // Unique email
	Password  string    `json:"password" pg:"password"`     // User password (hashed)
	Salt      string    `json:"salt" pg:"salt"`             // Salt used for password hashing
	FirstName string    `json:"first_name" pg:"first_name"` // User's first name
	LastName  string    `json:"last_name" pg:"last_name"`   // User's last name
	CreatedAt time.Time `json:"created_at" pg:"created_at"` // Date and time of user creation
	UpdatedAt time.Time `json:"updated_at" pg:"updated_at"` // Date and time of the last update
}

// BeforeInsert hook to set default UUID and generate a salt if not set
func (u *User) BeforeInsert() error {
	if u.ID == "" {
		newID := uuid.New().String()
		u.ID = newID
	}

	// Generate a random salt if not already set
	if u.Salt == "" {
		salt, err := auth.GenerateSalt()
		if err != nil {
			return err
		}
		u.Salt = salt
	}

	return nil
}

// TableName sets the table name for the User model
func (User) TableName() string {
	return "users"
}
