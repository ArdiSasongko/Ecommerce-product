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

func (s *CategoryService) getCategoriesWithPagination(ctx context.Context, params model.PaginatinParams) ([]model.CategoryResponse, error) {
	getRedis, err := s.c.Category.Get(ctx)
	if err != nil {
		return nil, err
	}

	if getRedis != nil {
		log.Println("get from redis")
		return ApplyPaginationCategoris(getRedis, params), nil
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

	return ApplyPaginationCategoris(resps, params), nil
}

func (s *CategoryService) GetCategories(ctx context.Context, params model.PaginatinParams) (*model.CategoryWithPaginationResponse, error) {
	data, err := s.getCategoriesWithPagination(ctx, params)
	if err != nil {
		return nil, err
	}

	totalCount := 0
	if getRedis, err := s.c.Category.Get(ctx); err == nil && getRedis != nil {
		totalCount = len(getRedis)
	} else {
		count, err := s.q.CountCategories(ctx)
		if err != nil {
			return nil, err
		}
		totalCount = int(count)
	}

	return &model.CategoryWithPaginationResponse{
		Data:       data,
		TotalCount: totalCount,
		Offset:     params.Offset,
		Limit:      params.Limit,
	}, nil
}
