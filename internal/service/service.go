package service

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-product/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/ArdiSasongko/Ecommerce-product/internal/storage/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Product interface {
		CreateProduct(context.Context, *model.ProductPayload) error
		UpdateProduct(context.Context, *model.ProductUpdatePayload) (*model.ProductUpdateResponse, error)
		UpdateVariant(context.Context, *model.VariantsUpdatePayload) (*model.VariantUpdateResponse, error)
	}
	Category interface {
		InsertCategory(context.Context, string) error
	}
}

func NewService(db *pgxpool.Pool, auth auth.JWTAuth) Service {
	q := sqlc.New(db)
	return Service{
		Product: &ProductService{
			q:  q,
			db: db,
		},
		Category: &CategoryService{
			q: q,
		},
	}
}
