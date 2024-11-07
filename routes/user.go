package routes

import (
	"github.com/drive-deep/auth-microservices/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes sets up routes related to user operations (e.g., fetch user details).
func SetupUserRoutes(app *fiber.App) {
	// GET route for fetching user details
	app.Get("/user", controllers.GetUserDetails)
}
