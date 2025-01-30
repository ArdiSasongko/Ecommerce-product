package handler

import (
	"strings"

	"github.com/ArdiSasongko/Ecommerce-product/internal/external"
	"github.com/gofiber/fiber/v2"
)

type MiddlewareHandler struct {
	external external.External
}

func (m *MiddlewareHandler) AdminMiddleware(accessLevel int32) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authToken := ctx.Get("Authorization")
		if authToken == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing token header authorization",
			})
		}

		rContext := ctx.Context()
		parts := strings.Split(authToken, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "header are malformed",
			})
		}

		token := parts[1]

		resp, err := m.external.User.Profile(rContext, token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		access := AccessLevel(accessLevel, resp.Data.Role)
		if !access {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "only highest level can access this",
			})
		}

		return ctx.Next()
	}
}

func AccessLevel(access, roleLevel int32) bool {
	return roleLevel >= access
}
