package middlewares

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/drive-deep/auth-microservices/auth"
	"github.com/gofiber/fiber/v2"
)

// Secret key for signing JWT (in a real application, store this securely)
var JWTSecret = []byte("your-secret-key")

// Token claims structure (you can add more fields as needed)
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func TokenAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header (bearer <token>)
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			// Missing token
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization token",
			})
		}

		// The token is expected to be in the form: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			// Token validation failed (invalid, expired, etc.)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Correctly access the claims from the MapClaims
		mapClaims := *claims // Dereference the pointer

		// Correctly access the claims from the MapClaims
		userID, _ := mapClaims["user_id"].(string)
		email, _ := mapClaims["email"].(string)

		// Store the claims in the context
		c.Locals("user_id", userID)
		c.Locals("email", email)

		// If the token is valid, pass the request to the next handler
		return c.Next()
	}
}
