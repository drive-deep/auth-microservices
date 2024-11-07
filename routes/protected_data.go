package routes

import (
	middlewares "github.com/drive-deep/auth-microservices/middleware"
	"github.com/gofiber/fiber/v2"
)

// ProtectedDataRoute defines the route for fetching protected data
func ProtectedDataRoute(app *fiber.App) {
	// Define the protected route with the middleware
	app.Get("/protected/secure-data", middlewares.TokenAuthMiddleware(), func(c *fiber.Ctx) error {
		// Retrieve user info from the context set by the middleware
		userID := c.Locals("user_id")
		email := c.Locals("email")

		// Return protected data
		return c.JSON(fiber.Map{
			"message": "This is protected data",
			"user_id": userID,
			"email":   email,
		})
	})
}
