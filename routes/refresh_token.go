package routes

import (
	"github.com/drive-deep/auth-microservices/auth"
	"github.com/gofiber/fiber/v2"
)

// RefreshTokenRoute defines the route to refresh the access token
func RefreshTokenRoute(app *fiber.App) {
	app.Post("/auth/refresh", func(c *fiber.Ctx) error {
		// Get the refresh token from the request body
		var requestBody struct {
			Token string `json:"token"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Call the RefreshToken function from the auth package
		newAccessToken, err := auth.RefreshToken(requestBody.Token, 1)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Return the new access token
		return c.JSON(fiber.Map{
			"refresh_token": newAccessToken,
		})
	})
}
