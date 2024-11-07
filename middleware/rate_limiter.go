package middlewares

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
)

var ctx = context.Background()

// RateLimitMiddleware is a middleware that limits the number of requests per minute
func RateLimitMiddleware(c *fiber.Ctx) error {
	// Set up Redis client (assumes Redis is running on localhost:6379)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Change this if Redis is hosted elsewhere
	})

	// Use the client's Get function to fetch the count of requests for the IP
	ip := c.IP()
	key := "rate_limit:" + ip
	ttl := 60 * time.Second // Time to live (1 minute)

	// Increment the request counter for this IP
	requestCount, err := client.Incr(ctx, key).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// If the count is 1, set the TTL (time-to-live) for the key
	if requestCount == 1 {
		client.Expire(ctx, key, ttl)
	}

	// If the request count exceeds the limit (e.g., 100 requests per minute), return an error
	if requestCount > 100 {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Rate limit exceeded. Try again later.",
		})
	}

	// Allow the request to proceed if rate limit is not exceeded
	return c.Next()
}
