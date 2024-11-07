package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" pg:"id,pk"`              // Primary key
	Email     string    `json:"email" pg:"email,unique"`    // Unique email
	Password  string    `json:"password" pg:"password"`     // User password
	FirstName string    `json:"first_name" pg:"first_name"` // User's first name
	LastName  string    `json:"last_name" pg:"last_name"`   // User's last name
	CreatedAt time.Time `json:"created_at" pg:"created_at"` // Date and time of user creation
	UpdatedAt time.Time `json:"updated_at" pg:"updated_at"` // Date and time of the last update
}

// TableName sets the table name for the User model (this is optional with go-pg if the struct name is the same)
func (User) TableName() string {
	return "users"
}
