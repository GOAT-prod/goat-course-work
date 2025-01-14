package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"search-service/settings"
)

func NewRedisClient(ctx context.Context, redisSettings settings.Redis) (*redis.Client, error) {
	options := &redis.Options{
		Addr:     redisSettings.Address,
		Password: redisSettings.Password,
		DB:       redisSettings.Database,
	}

	client := redis.NewClient(options)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
