package redis

import (
	"ddd-boilerplate/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(config *config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.Host,
		DB:   config.DB,
	})
}
