package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/ArdiSasongko/Ecommerce-product/internal/storage/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	q  *sqlc.Queries
	db *pgxpool.Pool
}

func (s *ProductService) insertProduct(ctx context.Context, qtx *sqlc.Queries, payload *model.ProductPayload) (int32, error) {
	priceStr := fmt.Sprintf("%.2f", payload.Price)
	priceNumeric := pgtype.Numeric{}
	if err := priceNumeric.Scan(priceStr); err != nil {
		return 0, err
	}

	resp, err := qtx.InsertProduct(ctx, sqlc.InsertProductParams{
		Name: payload.Name,
		Description: pgtype.Text{
			String: payload.Description,
			Valid:  true,
		},
		Price: priceNumeric,
	})
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (s *ProductService) insertVariant(ctx context.Context, qtx *sqlc.Queries, payload *model.VariantsPayload) (sqlc.InsertProductVariantRow, error) {
	resp, err := qtx.InsertProductVariant(ctx, sqlc.InsertProductVariantParams{
		ProductID: payload.ProductID,
		Color:     payload.Color,
		Size:      payload.Size,
		Quantity:  payload.Quantity,
	})

	if err != nil {
		return sqlc.InsertProductVariantRow{}, err
	}

	return resp, nil
}

func (s *ProductService) insertCategorProduct(ctx context.Context, qtx *sqlc.Queries, category string, id int32) error {
	cat, err := qtx.GetCategory(ctx, category)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("this %s category didnt exits", category)
		}
		return err
	}

	if err := qtx.InsertProductCategory(ctx, sqlc.InsertProductCategoryParams{
		ProductID:  id,
		CategoryID: cat.ID,
	}); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) CreateProduct(ctx context.Context, payload *model.ProductPayload) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := s.q.WithTx(tx)

	id, err := s.insertProduct(ctx, qtx, payload)
	if err != nil {
		return err
	}

	for _, variant := range payload.VariantProduct {
		variant.ProductID = id
		_, err := s.insertVariant(ctx, qtx, &variant)
		if err != nil {
			return err
		}
	}

	for _, cat := range payload.Categories {
		if err := s.insertCategorProduct(ctx, qtx, cat, id); err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
