package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

// InitRedis initializes the Redis connection
func InitRedis(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	Client = redis.NewClient(opts)

	// Test connection
	ctx := context.Background()
	if err := Client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	log.Println("Redis connected successfully")
	return Client, nil
}

// GetClient returns the Redis client instance
func GetClient() *redis.Client {
	return Client
}
