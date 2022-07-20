package redis

import "github.com/go-redis/redis/v8"

type RDB struct {
	DB *redis.Client
}
