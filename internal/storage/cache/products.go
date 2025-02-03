package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/redis/go-redis/v9"
)

type ProductCache struct {
	client *redis.Client
}

func (r *ProductCache) SetProducts(ctx context.Context, payload []model.ProductsResponse) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshalling data :%w", err)
	}

	if err := r.client.SetEx(ctx, ProductsKey, data, MaxSet).Err(); err != nil {
		return fmt.Errorf("failed set to redis :%v", err)
	}

	return nil
}

func (r *ProductCache) GetProducts(ctx context.Context) ([]model.ProductsResponse, error) {
	data, err := r.client.Get(ctx, ProductsKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed get data from redis :%v", err)
	}

	var categories []model.ProductsResponse
	if data != "" {
		if err := json.Unmarshal([]byte(data), &categories); err != nil {
			return nil, fmt.Errorf("failed to unmarshaling data :%w", err)
		}
	}

	return categories, err
}
