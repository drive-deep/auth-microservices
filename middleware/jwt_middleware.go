package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JWTMiddleware is a middleware to validate JWT token from the Authorization header
func JWTMiddleware(c *fiber.Ctx) error {
	// Get the token from the "Authorization" header
	tokenString := c.Get("Authorization")

	// Remove "Bearer " from the beginning of the string if it exists
	if len(tokenString) > 7 && strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or malformed token",
		})
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		// Return the JWT secret key to validate the signature
		return []byte("your_jwt_secret_key"), nil
	})

	// If there's an error parsing the token, return an unauthorized error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Token is valid, now check the claims (e.g., expiration time)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the "exp" field from the claims
		expirationTime := claims["exp"].(float64)
		// If the token is expired, return an unauthorized error
		if time.Now().Unix() > int64(expirationTime) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token has expired",
			})
		}

		// Add the user's email to the context, you can use it in your handlers
		c.Locals("email", claims["email"])

		// Proceed to the next handler in the stack
		return c.Next()
	}

	// If we get here, the token is invalid
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid token",
	})
}
