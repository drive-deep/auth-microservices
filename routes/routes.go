package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes centralizes all the route setups
func SetupRoutes(app *fiber.App) {
	// Setup authentication routes
	SetupAuthRoutes(app)

	// Setup user-related routes
	SetupUserRoutes(app)
}
