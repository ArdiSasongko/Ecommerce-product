package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/redis/go-redis/v9"
)

type CategoryCache struct {
	client *redis.Client
}

func (r *CategoryCache) Set(ctx context.Context, payload []model.CategoryResponse) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshalling data :%w", err)
	}

	if err := r.client.SetEx(ctx, CategoriesKey, data, MaxSet).Err(); err != nil {
		return fmt.Errorf("failed set to redis :%v", err)
	}

	return nil
}

func (r *CategoryCache) Get(ctx context.Context) ([]model.CategoryResponse, error) {
	data, err := r.client.Get(ctx, CategoriesKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed get data from redis :%v", err)
	}

	var categories []model.CategoryResponse
	if data != "" {
		if err := json.Unmarshal([]byte(data), &categories); err != nil {
			return nil, fmt.Errorf("failed to unmarshaling data :%w", err)
		}
	}

	return categories, err
}
