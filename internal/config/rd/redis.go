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

	ctx := context.Background()
	ping, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	log.Info("PING redis: " + ping)
	return client
}
