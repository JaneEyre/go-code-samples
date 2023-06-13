package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func ping(client *redis.Client) error {
	// For the demo, we need only a background context
	ctx := context.Background()
	// Ping the redis server. It should respond with "PONG".
	fmt.Println(client.Ping(ctx))

	// Get the client info.
	info, err := client.ClientInfo(ctx).Result()
	if err != nil {
		return fmt.Errorf("method ClientInfo failed: %w", err)
	}

	fmt.Printf("%#v\n", info)
	return nil
}
