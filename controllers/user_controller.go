package controllers

import (
	"net/http"

	"github.com/drive-deep/auth-microservices/config"
	"github.com/drive-deep/auth-microservices/models"
	"github.com/gofiber/fiber/v2"
)

// GetUserDetails handles the route to fetch user details
func GetUserDetails(c *fiber.Ctx) error {
	// Retrieve the email of the user from the JWT token (using middleware)
	email := c.Locals("email").(string)

	// Query the user from the database by their email using go-pg
	var user models.User
	err := config.DB.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Return the user details
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id":    user.ID,
		"email": user.Email,
	})
}
