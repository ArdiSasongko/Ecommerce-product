package rd

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-product/internal/config/logger"
	"github.com/redis/go-redis/v9"
)

var log = logger.NewLogger()

func NewRedis(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to connected redis :%v", err)
	}

	log.Info("PING redis: " + ping)
	return client
}
