package cache

import (
	"context"
	"time"

	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/redis/go-redis/v9"
)

const (
	CategoriesKey = "ListCategories"
	ProductsKey   = "ListProducts"
	MaxSet        = 10 * time.Minute
)

type RedisCache struct {
	Product interface {
		SetProducts(context.Context, []model.ProductsResponse) error
		GetProducts(context.Context) ([]model.ProductsResponse, error)
	}
	Category interface {
		Set(context.Context, []model.CategoryResponse) error
		Get(context.Context) ([]model.CategoryResponse, error)
	}
}

func NewRedisCache(client *redis.Client) RedisCache {
	return RedisCache{
		Product: &ProductCache{
			client: client,
		},
		Category: &CategoryCache{
			client: client,
		},
	}
}
