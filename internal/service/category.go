package service

import (
	"context"
	"log"

	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/ArdiSasongko/Ecommerce-product/internal/storage/cache"
	"github.com/ArdiSasongko/Ecommerce-product/internal/storage/sqlc"
)

type CategoryService struct {
	q *sqlc.Queries
	c cache.RedisCache
}

func (s *CategoryService) InsertCategory(ctx context.Context, name string) error {
	_, err := s.q.CreateCategory(ctx, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, name string, paramName string) (string, error) {
	resp, err := s.q.UpdateCategory(ctx, sqlc.UpdateCategoryParams{
		Name:   name,
		Name_2: paramName,
	})
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, name string) error {
	if err := s.q.DeleteCategory(ctx, name); err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) GetCategory(ctx context.Context) ([]model.CategoryResponse, error) {
	getRedis, err := s.c.Category.Get(ctx)
	if err != nil {
		return nil, err
	}

	if getRedis != nil {
		log.Println("get from redis")
		return getRedis, nil
	}

	categories, err := s.q.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	var resps []model.CategoryResponse
	for _, category := range categories {
		resp := model.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt.Time,
		}
		resps = append(resps, resp)
	}

	log.Println("get from db")

	if err := s.c.Category.Set(ctx, resps); err != nil {
		log.Printf("failed to set data in Redis: %v\n", err)
	}

	return resps, nil
}
