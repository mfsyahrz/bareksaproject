package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mfsyahrz/bareksaproject/internal/config"
)

type Cache interface {
	Set(ctx context.Context, key string, val interface{}) error
	Get(ctx context.Context, key string, val interface{}) error
	Del(ctx context.Context, key string) error
}

type cache struct {
	client *redis.Client
	config *config.Redis
}

func New(config *config.Redis) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err.Error())
	}

	return &cache{rdb, config}
}

func (c *cache) Set(ctx context.Context, key string, val interface{}) error {
	return c.client.Set(ctx, key, val, time.Minute*time.Duration(c.config.DefaultTTL)).Err()
}

func (c *cache) Get(ctx context.Context, key string, val interface{}) error {
	return c.client.Get(ctx, key).Scan(val)
}

func (c *cache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
