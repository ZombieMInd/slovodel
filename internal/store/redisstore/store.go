package redisstore

import (
	"github.com/go-redis/redis"
)

type Store struct {
	client *redis.Client
}

func New(c *redis.Client) *Store {
	return &Store{client: c}
}
