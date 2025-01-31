package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/ArdiSasongko/Ecommerce-product/internal/storage/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

const CtxTimeout = time.Second * 5

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
	ctx, cancel := context.WithTimeout(ctx, CtxTimeout)
	defer cancel()

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

func (s *ProductService) UpdateProduct(ctx context.Context, payload *model.ProductUpdatePayload) (*model.ProductUpdateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeout)
	defer cancel()

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := s.q.WithTx(tx)

	resp, err := qtx.GetProduct(ctx, payload.ProductID)
	if err != nil {
		return nil, err
	}

	var prod sqlc.UpdateProductParams

	if payload.Name != nil {
		prod.Name = *payload.Name
	} else {
		prod.Name = resp.Name
	}

	if payload.Description != nil {
		prod.Description = pgtype.Text{
			String: *payload.Description,
			Valid:  true,
		}
	} else {
		prod.Description = resp.Description
	}

	if payload.Price != nil {
		priceStr := fmt.Sprintf("%.2f", *payload.Price)
		priceNumeric := pgtype.Numeric{}
		if err := priceNumeric.Scan(priceStr); err != nil {
			return nil, err
		}
		prod.Price = priceNumeric
	} else {
		prod.Price = resp.Price
	}

	prod.ID = payload.ProductID

	newProd, err := qtx.UpdateProduct(ctx, prod)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	priceFloat, _ := newProd.Price.Float64Value()
	return &model.ProductUpdateResponse{
		Name:        newProd.Name,
		Description: newProd.Description.String,
		Price:       float32(priceFloat.Float64),
		UpdateAt:    newProd.UpdatedAt.Time,
	}, nil
}

func (s *ProductService) UpdateVariant(ctx context.Context, payload *model.VariantsUpdatePayload) (*model.VariantUpdateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeout)
	defer cancel()

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := s.q.WithTx(tx)

	resp, err := qtx.GetVariantByID(ctx, sqlc.GetVariantByIDParams{
		ID:        payload.VariantID,
		ProductID: payload.ProductID,
	})
	if err != nil {
		return nil, err
	}

	var vars sqlc.UpdateProductVariantParams

	if payload.Color != nil {
		vars.Color = *payload.Color
	} else {
		vars.Color = resp.Color
	}

	if payload.Size != nil {
		vars.Size = *payload.Size
	} else {
		vars.Size = resp.Size
	}

	if payload.Quantity != nil {
		vars.Quantity = *payload.Quantity
	} else {
		vars.Quantity = resp.Quantity
	}

	vars.ID = payload.VariantID
	vars.ProductID = payload.ProductID

	newVars, err := qtx.UpdateProductVariant(ctx, vars)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &model.VariantUpdateResponse{
		Color:     newVars.Color,
		Size:      newVars.Size,
		Quantity:  newVars.Quantity,
		UpdatedAt: newVars.UpdatedAt.Time,
	}, nil
}

func (s *ProductService) deleteVariants(ctx context.Context, qtx *sqlc.Queries, productId, id int32) error {
	if err := qtx.DeleteProductVariant(ctx, sqlc.DeleteProductVariantParams{
		ProductID: productId,
		ID:        id,
	}); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int32) error {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeout)
	defer cancel()

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := s.q.WithTx(tx)
	resp, err := qtx.GetVariantsByProductID(ctx, id)
	if err != nil {
		return nil
	}

	for _, res := range resp {
		if err := s.deleteVariants(ctx, qtx, id, res.ID); err != nil {
			return err
		}
	}

	if err := qtx.DeleteProduct(ctx, id); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
