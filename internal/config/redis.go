package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(ctx context.Context, addr string) (*RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisRepository{client: client}, nil
}
