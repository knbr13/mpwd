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
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading '.env' file: %v", err)
	}

	// Parse Upstash Redis URL from environment variable
	redisURL := os.Getenv("UPSTASH_REDIS_REST_URL")
	if redisURL == "" {
		log.Fatal("UPSTASH_REDIS_REST_URL environment variable is not set")
	}

	// Create a context
	ctx := context.Background()

	// Create Redis client
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Error parsing Redis URL: %v", err)
	}
	client := redis.NewClient(opt)
	defer client.Close()

	// Check connection to Redis
	res, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	fmt.Printf("Ping: %s\n", res)

	// Set key-value pairs for different data types
	// The last argument is the expiration, zero expiration means the key has no expiration time
	client.Set(ctx, "string", "hello world", 0)
	client.Set(ctx, "int", 42, 0)
	client.Set(ctx, "float", 3.14, 0)
	client.Set(ctx, "bool", true, 0)

	// Retrieve and print values for different data types
	stringValue, err := client.Get(ctx, "string").Result()
	if err != nil {
		log.Fatalf("Error getting value for key 'string': %v", err)
	}
	fmt.Printf("String Value: %q\n", stringValue)

	intValue, err := client.Get(ctx, "int").Int()
	if err != nil {
		log.Fatalf("Error getting value for key 'int': %v", err)
	}
	fmt.Printf("Integer Value: %d\n", intValue)

	floatValue, err := client.Get(ctx, "float").Float64()
	if err != nil {
		log.Fatalf("Error getting value for key 'float': %v", err)
	}
	fmt.Printf("Float Value: %f\n", floatValue)

	boolValue, err := client.Get(ctx, "bool").Bool()
	if err != nil {
		log.Fatalf("Error getting value for key 'bool': %v", err)
	}
	fmt.Printf("Boolean Value: %t\n", boolValue)

	// Use Set to update a value of a key
	client.Set(ctx, "string", "hello world updated", 0)
	stringValue, err = client.Get(ctx, "string").Result()
	if err != nil {
		log.Fatalf("Error getting value for key 'string': %v", err)
	}
	fmt.Printf("String Value after update: %q\n", stringValue)

	// Key to set then delete
	key := "to_be_deleted"
	client.Set(ctx, key, "foo", 0)

	// Delete the key
	err = client.Del(ctx, key).Err()
	if err != nil {
		log.Fatalf("Error deleting key %q: %v", key, err)
	}

	// Check if the key has been successfully deleted
	_, err = client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("Key %q deleted successfully!\n", key)
	} else if err != nil {
		log.Fatalf("Error checking key %q after deletion: %v", key, err)
	} else {
		fmt.Printf("Key %q still exists after deletion!\n", key)
	}
}
