package handler

import (
	"github.com/ArdiSasongko/Ecommerce-product/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-product/internal/external"
	"github.com/ArdiSasongko/Ecommerce-product/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Health interface {
		Check(*fiber.Ctx) error
	}
	Product interface {
		CreateProduct(*fiber.Ctx) error
	}
	Category interface {
		CreateCategory(*fiber.Ctx) error
	}
	Middleware interface {
		AdminMiddleware(int32) fiber.Handler
	}
}

func NewHandler(db *pgxpool.Pool, auth auth.JWTAuth) Handler {
	service := service.NewService(db, auth)
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
