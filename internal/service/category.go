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
