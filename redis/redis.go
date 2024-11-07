package redis

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type RedisClient struct {
	client *redis.Client
}

// NewRedisClient initializes a new Redis client
func NewRedisClient(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis server address (e.g., "localhost:6379")
		Password: password, // Redis password
		DB:       db,       // Redis DB index
	})

	return &RedisClient{client: rdb}
}

// Close closes the Redis connection
func (r *RedisClient) Close() error {
	return r.client.Close()
}

// Set sets a key-value pair in Redis with an expiry datetime (expiration timestamp)
// It will only set the key if the expiration time is in the future
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expirationTime time.Time) error {
	// Check if the expiration time is in the future
	if expirationTime.Before(time.Now()) {
		return fmt.Errorf("cannot set key, expiration time is in the past")
	}

	// Set the key with the expiration timestamp
	err := r.client.SetEX(ctx, key, value, time.Until(expirationTime)).Err()
	if err != nil {
		return fmt.Errorf("could not set value in redis: %v", err)
	}

	return nil
}

// Get retrieves a value from Redis by key
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	} else if err != nil {
		return "", fmt.Errorf("could not get value from redis: %v", err)
	}
	return val, nil
}

// Delete removes a key-value pair from Redis
func (r *RedisClient) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("could not delete key from redis: %v", err)
	}
	return nil
}

func (r *RedisClient) Reconnect(ctx context.Context) error {
	maxRetries := 5
	retryInterval := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempting to reconnect to Redis, attempt %d/%d", i+1, maxRetries)
		// Try to ping the Redis server
		_, err := r.client.Ping(ctx).Result()
		if err == nil {
			log.Println("Reconnected to Redis successfully")
			return nil
		}

		// Wait for a while before retrying
		log.Printf("Redis connection failed, retrying in %v...", retryInterval)
		time.Sleep(retryInterval)
	}

	return fmt.Errorf("failed to reconnect to Redis after %d attempts", maxRetries)
}

// AutoReconnect is a background goroutine that will periodically check the Redis connection
// and attempt to reconnect if the connection is lost
func (r *RedisClient) AutoReconnect(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second) // Check every 10 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("AutoReconnect stopping...")
			return
		case <-ticker.C:
			// Check if Redis is still connected by pinging it
			_, err := r.client.Ping(ctx).Result()
			if err != nil {
				// If connection is lost, attempt to reconnect
				log.Printf("Lost connection to Redis, attempting to reconnect...")
				if reconnectErr := r.Reconnect(ctx); reconnectErr != nil {
					log.Printf("Failed to reconnect to Redis: %v", reconnectErr)
				}
			}
		}
	}
}

// InitRedis initializes the Redis client and checks connectivity
func InitRedis() (*RedisClient, error) {
	// Retrieve Redis credentials from environment variables
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0 // Default to DB index 0, modify if needed

	// Initialize the Redis client
	client := NewRedisClient(redisAddr, redisPassword, redisDB)

	// Ping the Redis server to check connectivity
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")
	return client, nil
}
