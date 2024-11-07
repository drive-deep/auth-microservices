package routes

import (
	"github.com/drive-deep/auth-microservices/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes sets up routes related to authentication (signup, login).
func SetupAuthRoutes(app *fiber.App) {
	// POST route for user signup
	app.Post("/signup", controllers.SignUp)

	// POST route for user login
	app.Post("/login", controllers.Login)
}
