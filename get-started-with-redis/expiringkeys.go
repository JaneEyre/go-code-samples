package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func expiringKeys(client *redis.Client) error {
	ctx := context.Background()

	// Add a temporary player
	err := client.HSet(ctx, "player:10", "name", "Crymyios", "score", 0, "team", "Knucklewimp", "challenges_completed", 0).Err()
	if err != nil {
		return fmt.Errorf("cannot set player:10: %w", err)
	}

	// Set an expiration time for player:10
	if !client.Expire(ctx, "player:10", time.Second).Val() {
		return fmt.Errorf("cannot set expiration time for player:10")
	}

	// Get player:10
	for i := 0; i < 3; i++ {
		val, err := client.HGet(ctx, "player:10", "name").Result()
		if err != nil {
			fmt.Printf("player:10 has expired: %v\n", err)
			return nil
		}
		fmt.Printf("player:10's name: %s\n", val)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}
