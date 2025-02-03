package handler

import (
	"github.com/ArdiSasongko/Ecommerce-product/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-product/internal/external"
	"github.com/ArdiSasongko/Ecommerce-product/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

const (
	defaultLimit = 5
	maxLimit     = 100
	minOffset    = 0
)

type Handler struct {
	Health interface {
		Check(*fiber.Ctx) error
	}
	Product interface {
		CreateProduct(*fiber.Ctx) error
		UpdateProduct(*fiber.Ctx) error
		UpdateVariant(*fiber.Ctx) error
		DeleteProduct(*fiber.Ctx) error
		GetProducts(*fiber.Ctx) error
		GetProduct(*fiber.Ctx) error
	}
	Category interface {
		CreateCategory(*fiber.Ctx) error
		UpdateCategory(*fiber.Ctx) error
		DeleteCategory(*fiber.Ctx) error
		GetCategories(*fiber.Ctx) error
	}
	Middleware interface {
		AdminMiddleware(int32) fiber.Handler
	}
}

func NewHandler(db *pgxpool.Pool, auth auth.JWTAuth, cache *redis.Client) Handler {
	service := service.NewService(db, auth, cache)
	external := external.NewExternal()
	return Handler{
		Health: &HealthHandler{},
		Product: &ProductHandler{
			service: service,
		},
		Category: &CategoryHandler{
			service: service,
		},
		Middleware: &MiddlewareHandler{
			external: external,
		},
	}
}
