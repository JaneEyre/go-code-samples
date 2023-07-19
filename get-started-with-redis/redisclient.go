package main

import "github.com/redis/go-redis/v9"

func newClient(conn string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conn,
		DB:       db,
		Password: "",
	})
}
