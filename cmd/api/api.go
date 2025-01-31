package api

import (
	"github.com/ArdiSasongko/Ecommerce-product/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Application struct {
	config  Config
	handler handler.Handler
}

type Config struct {
	addrHTTP string
	log      *logrus.Logger
	db       DBConfig
	auth     AuthConfig
	redis    RedisConfig
}

type DBConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type AuthConfig struct {
	secret string
	iss    string
	aud    string
}

type RedisConfig struct {
	addr string
}

func (a *Application) Mount() *fiber.App {
	r := fiber.New()

	r.Get("/health", a.handler.Health.Check)

	v1 := r.Group("/v1")

	product := v1.Group("/products")
	product.Post("/", a.handler.Middleware.AdminMiddleware(2), a.handler.Product.CreateProduct)
	product.Patch("/:productID", a.handler.Middleware.AdminMiddleware(2), a.handler.Product.UpdateProduct)
	product.Patch("/:productID/variants/:variantID", a.handler.Middleware.AdminMiddleware(2), a.handler.Product.UpdateVariant)

	category := v1.Group("/category")
	category.Post("/", a.handler.Middleware.AdminMiddleware(2), a.handler.Category.CreateCategory)
	category.Put("/:category_name", a.handler.Middleware.AdminMiddleware(2), a.handler.Category.UpdateCategory)
	return r
}

func (a *Application) Run(r *fiber.App) error {
	a.config.log.Printf("http server has run, port%v", a.config.addrHTTP)
	return r.Listen(a.config.addrHTTP)
}
