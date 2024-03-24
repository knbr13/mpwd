package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Create a context
	ctx := context.Background()

	// Parse Upstash Redis URL from environment variable
	redisURL := os.Getenv("UPSTASH_REDIS_REST_URL")
	if redisURL == "" {
		log.Fatal("UPSTASH_REDIS_REST_URL environment variable is not set")
	}

	// Create Redis client
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Error parsing Redis URL: %v", err)
	}
	client := redis.NewClient(opt)
	defer client.Close()

	// Set key-value pair
	key := "hello"
	err = client.Set(ctx, key, "world", 0).Err()
	if err != nil {
		log.Fatalf("Error setting key-value pair: %v", err)
	}

	// Get value for the key
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Error getting value for key: %v", err)
	}
	fmt.Printf("Key: %q | Value: %q\n", key, value)
}
