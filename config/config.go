package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

// DB holds the database connection instance
var DB *pg.DB

// InitDB initializes the database connection using go-pg
func InitDB() {
	// Load environment variables from .env file

	// Retrieve the PostgreSQL credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Establish the connection to the PostgreSQL database using go-pg
	DB = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})

	// Ping the database to check the connection
	_, errDb := DB.Exec("SELECT 1")
	if errDb != nil {
		log.Fatal("Failed to connect to the database:", errDb)
	}

	log.Println("Successfully connected to the database")
}

// SetupAppConfig sets up the application-wide configurations like middleware, logging, etc.
func SetupAppConfig(app *fiber.App) {
	// Set up global middleware or app settings here
	app.Use(func(c *fiber.Ctx) error {
		// Log the request method and path
		log.Printf("Received %s request on %s", c.Method(), c.Path())
		return c.Next()
	})

	// You can add more global middleware, such as authentication, error handling, etc.
}
