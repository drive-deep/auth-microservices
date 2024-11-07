package controllers

import (
	"log"
	"net/http"

	"github.com/drive-deep/auth-microservices/config" // Update this path to your actual project path
	"github.com/drive-deep/auth-microservices/models" // Update this path to your actual project path

	"github.com/gofiber/fiber/v2"
)

// GetUserDetails handles GET requests to fetch all user details
// GetUserDetails handles GET requests to fetch specific user details
func GetUserDetails(c *fiber.Ctx) error {
	// Define a slice to hold the list of users with only specific fields
	var users []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	// Fetch specific fields (first_name, last_name, email) from the database
	err := config.DB.Model(&models.User{}).
		Column("first_name", "last_name", "email"). // go-pg's Column method to select specific fields
		Select(&users)                              // Use Select to fill the slice with the data
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	if len(users) == 0 {
		return c.Status(http.StatusOK).JSON([]interface{}{})
	}
	// Return the list of users in JSON format
	return c.Status(http.StatusOK).JSON(users)
}
