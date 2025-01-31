package service

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-product/internal/storage/sqlc"
)

type CategoryService struct {
	q *sqlc.Queries
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
