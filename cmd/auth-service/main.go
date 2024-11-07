package main

import (
	"log"

	"github.com/drive-deep/auth-microservices/config"
	"github.com/drive-deep/auth-microservices/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize database connection

	config.InitDB()

	// Create a new Fiber app
	app := fiber.New()

	// Apply global middlewares
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n", // Customize log format if needed
		TimeFormat: "02-Jan-2006",
	})) // Logging middleware for request logging
	//app.Use(middlewares.RateLimitMiddleware) // Apply rate limiting

	// Define routes and handlers
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	// Set up routes
	routes.SetupRoutes(app)

	// Start the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
