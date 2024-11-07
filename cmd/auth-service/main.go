package main

import (
	"log"

	"github.com/drive-deep/auth-microservices/config"
	middlewares "github.com/drive-deep/auth-microservices/middleware"
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
	app.Use(logger.New())                    // Logging middleware for request logging
	app.Use(middlewares.RateLimitMiddleware) // Apply rate limiting

	// Set up routes
	routes.SetupRoutes(app)

	// Start the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
