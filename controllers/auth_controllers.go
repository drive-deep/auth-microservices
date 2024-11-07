package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/drive-deep/auth-microservices/auth"
	"github.com/drive-deep/auth-microservices/config"
	"github.com/drive-deep/auth-microservices/models"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// SignUp handles user sign-up
// SignUp handles user sign-up
func SignUp(c *fiber.Ctx) error {
	// Parse the request body
	var req SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	// Validate input (you can add more validation if needed)
	if req.Email == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	// Check if the email already exists in the database using CheckEmailExists
	emailExists, err := CheckEmailExists(config.DB, req.Email)
	if err != nil {
		log.Printf("Error checking email existence: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// If the email already exists, return a 400 Bad Request with the error message
	if emailExists {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Email already exists",
			"request": req, // Include the request body in the response
		})
	}

	// Proceed with sign-up since the email does not exist
	// Generate a salt
	salt, err := auth.GenerateSalt()
	if err != nil {
		log.Printf("Error generating salt: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// Hash the password with the salt
	hashedPassword, err := auth.HashPasswordWithSalt(req.Password, salt)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	fmt.Println(req.Password)
	fmt.Println(salt)
	fmt.Println(hashedPassword)
	fmt.Println(string(hashedPassword))

	// Create a new user object
	user := models.User{
		ID:        uuid.New().String(),
		Email:     req.Email,
		Password:  string(hashedPassword), // Store hashed password
		Salt:      salt,                   // Store salt
		FirstName: req.FirstName,
		LastName:  req.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the user to the database
	_, err = config.DB.Model(&user).Insert()
	if err != nil {
		log.Printf("Error inserting user into database: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Return success response
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    req,
	})
}

// LoginRequest struct to capture login details
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login handles user login and JWT token generation
func Login(c *fiber.Ctx) error {
	var req LoginRequest

	// Parse the request body
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	// Check if the user exists using CheckEmailExists
	emailExists, err := CheckEmailExists(config.DB, req.Email)
	if err != nil {
		log.Printf("Error checking email existence: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// If the user does not exist, return a 400 error
	if !emailExists {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "User not registered",
		})
	}

	// Retrieve the user from the database
	var user models.User
	err = config.DB.Model(&user).Where("email = ?", req.Email).Select()
	if err != nil {
		log.Printf("Error querying user: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// Hash the provided password with the salt stored in the DB
	hashedPassword, err := auth.HashPasswordWithSalt(req.Password, user.Salt)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating hashed password",
		})
	}
	
	// Compare the generated hash with the stored password hash
	if string(hashedPassword) != user.Password {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Email, 1)
	if err != nil {
		log.Printf("Error generating JWT token: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating token",
		})
	}

	// Return the JWT token
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// Update CheckEmailExists function to use the correct version of pg.DB
func CheckEmailExists(db *pg.DB, email string) (bool, error) {
	var user models.User
	err := db.Model(&user).Where("email = ?", email).Select()

	// Check if no rows were found
	if err == pg.ErrNoRows {
		return false, nil // Email does not exist
	} else if err != nil {
		// Database error other than no rows found
		return false, err
	}

	// If no error, it means the email already exists
	return true, nil
}
