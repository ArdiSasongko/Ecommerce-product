package cache

import (
	"github.com/redis/go-redis/v9"
)

type ProductCache struct {
	client *redis.Client
}
