package controllers

import (
	"log"
	"net/http"

	"github.com/drive-deep/auth-microservices/config" // Update this path to your actual project path
	"github.com/drive-deep/auth-microservices/models" // Update this path to your actual project path

	"github.com/gofiber/fiber/v2"
)

// GetUserDetails handles GET requests to fetch all user details
func GetUserDetails(c *fiber.Ctx) error {
	// Define a slice to hold the list of users
	var users []models.User

	// Fetch all users from the database
	err := config.DB.Model(&users).Select()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	// Return the list of users in JSON format
	return c.Status(http.StatusOK).JSON(users)
}
