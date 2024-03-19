package db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Datastore is an interface that our Database will implement.
type Datastore interface {
	IncrementCounter(ctx context.Context, key string) (int64, error)
}

// RedisDB wraps the Redis client.
type RedisDB struct {
	client *redis.Client
}

// NewRedisDB creates a new Database instance with a Redis client.
func NewRedisDB(addr string) *RedisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr, // Redis server address
		Password: "",   // no password set
		DB:       0,    // use default DB
	})
	return &RedisDB{client: rdb}
}

// IncrementCounter increments a number in Redis and returns the updated value.
func (db *RedisDB) IncrementCounter(ctx context.Context, key string) (int64, error) {
	return db.client.Incr(ctx, key).Result()
}
