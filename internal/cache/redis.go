package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisCache{
		client: client,
	}
}

type CachedResponse struct {
	StatusCode int
	Header     map[string][]string
	Body       []byte
}

var ctx = context.Background()

func (c *RedisCache) Get(key string) (CachedResponse, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return CachedResponse{}, fmt.Errorf("failed to get key %s from redis: %w\n", key, err)
	}
	var resp CachedResponse
	if err = json.Unmarshal([]byte(val), &resp); err != nil {
		return CachedResponse{}, fmt.Errorf("failed to unmarshal json to CachedResponse: %w\n", err)
	}
	return resp, nil
}

func (c *RedisCache) Set(key string, val CachedResponse) error {
	valJSON, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("failed to marshal CachedResponse to json: %w\n", err)
	}
	return c.client.Set(ctx, key, valJSON, 0).Err()
}

func (c *RedisCache) Clear() error {
	if err := c.client.FlushDB(ctx).Err(); err != nil {
		return err
	}
	return nil
}
