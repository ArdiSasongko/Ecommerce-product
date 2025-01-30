package cache

import "github.com/redis/go-redis/v9"

type RedisCache struct {
	User interface{}
}

func NewRedisCache(client *redis.Client) RedisCache {
	return RedisCache{}
}
