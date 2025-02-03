package service

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-product/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/ArdiSasongko/Ecommerce-product/internal/storage/cache"
	"github.com/ArdiSasongko/Ecommerce-product/internal/storage/sqlc"
	"github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Product interface {
		CreateProduct(context.Context, *model.ProductPayload) error
		UpdateProduct(context.Context, *model.ProductUpdatePayload) (*model.ProductUpdateResponse, error)
		UpdateVariant(context.Context, *model.VariantsUpdatePayload) (*model.VariantUpdateResponse, error)
		DeleteProduct(context.Context, int32) error
	}
	Category interface {
		InsertCategory(context.Context, string) error
		UpdateCategory(context.Context, string, string) (string, error)
		DeleteCategory(context.Context, string) error
		GetCategories(ctx context.Context, params model.PaginatinParams) (*model.CategoryWithPaginationResponse, error)
	}
}

func NewService(db *pgxpool.Pool, auth auth.JWTAuth, rd *redis.Client) Service {
	q := sqlc.New(db)
	cache := cache.NewRedisCache(rd)
	return Service{
		Product: &ProductService{
			q:  q,
			db: db,
		},
		Category: &CategoryService{
			q: q,
			c: cache,
		},
	}
}
