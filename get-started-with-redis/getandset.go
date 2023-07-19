package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Task: The quest is a lowercase string. Change it to title case.
// Then print the challenges in the order they have to be completed.

func getAndSet(client *redis.Client) error {
	ctx := context.Background()

	quest, err := client.Get(ctx, "quest").Result()
	if err != nil {
		return fmt.Errorf("cannot get quest: %w", err)
	}

	quest = cases.Title(language.English).String(quest)

	err = client.Set(ctx, "quest", quest, 0).Err()
	if err != nil {
		return fmt.Errorf("cannot update quest: %w", err)
	}

	fmt.Printf("Quest is now: %s\n", client.Get(ctx, "quest").Val())

	return nil
}
