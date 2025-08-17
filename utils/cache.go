package utils

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

// RedisCache represents a Redis cache client
// with methods for storing and retrieving cached data.
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache returns a new instance of RedisCache.
func NewRedisCache(addr string) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		})

	return &RedisCache{client: client}
}

// Store stores data in the Redis cache with the given key and expiration time.
func (r *RedisCache) Store(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, jsonData, expiration).Err()
}

// Get retrieves cached data from Redis by the given key.
func (r *RedisCache) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

// Delete removes the cached data with the given key from Redis.
func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Ping checks the connection to the Redis server.
func (r *RedisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func main() {}
