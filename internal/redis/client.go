package redis

import (
	"context"
	"fmt"
	"sandbox-gql/internal/env"

	"github.com/go-redis/redis"
)

func Client(ctx context.Context, redisVars *env.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisVars.Host(), redisVars.Port()),
		Password: redisVars.Password(),
		DB:       redisVars.DB(),
	})
	/*
		pong, err := client.Ping(ctx).Result()
		if err != nil {
			fmt.Println("Error connecting to Redis:", err)
		} else {
			fmt.Println("Connected to Redis:", pong)
		}*/

	return client, nil
}
