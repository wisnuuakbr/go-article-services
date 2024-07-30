package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Client *redis.Client
}

// function NewRedisCache is a Constructor
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{Client: client}
}

func (r *RedisCache) SetCache(key string, value interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(context.Background(), key, jsonData, expiration).Err()
}

func (r *RedisCache) GetCache(key string, dest interface{}) error {
	data, err := r.Client.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

func (r *RedisCache) DeleteCache(key string) error {
    return r.Client.Del(context.Background(), key).Err()
}