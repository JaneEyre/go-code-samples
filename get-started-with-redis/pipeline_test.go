package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func incrementScorePipe(client *redis.Client, player string) error {
	ctx := context.Background()
	client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 1000; i++ {
			err := pipe.HIncrBy(ctx, player, "score", 1).Err()
			if err != nil {
				return fmt.Errorf("cannot increment score for player %s to %d: %w", player, i, err)
			}
		}
		pipe.HSet(ctx, player, "score", 1)
		return nil
	})
	return nil
}

func incrementScoreNoPipe(client *redis.Client, player string) error {
	ctx := context.Background()
	for i := 0; i < 1000; i++ {
		err := client.HIncrBy(ctx, player, "score", 1).Err()
		if err != nil {
			return fmt.Errorf("cannot increment score for player %s to %d: %w", player, i, err)
		}
	}
	client.HSet(ctx, player, "score", 1)
	return nil
}

func BenchmarkPipeline(b *testing.B) {
	client := newClient(dbconn, 0)
	b.Run("PipelinedHIncrBy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			incrementScorePipe(client, "player:1")
		}
	})
}

func BenchmarkNoPipeline(b *testing.B) {
	client := newClient(dbconn, 0)
	b.Run("HIncrBy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			incrementScoreNoPipe(client, "player:2")
		}
	})
}
