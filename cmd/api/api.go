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

	return r
}

func (a *Application) Run(r *fiber.App) error {
	a.config.log.Printf("http server has run, port%v", a.config.addrHTTP)
	return r.Listen(a.config.addrHTTP)
}
