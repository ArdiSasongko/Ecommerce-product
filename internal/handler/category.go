package handler

import (
	"net/url"
	"strconv"

	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/ArdiSasongko/Ecommerce-product/internal/service"
	"github.com/ArdiSasongko/Ecommerce-product/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service service.Service
}

func (h *CategoryHandler) CreateCategory(ctx *fiber.Ctx) error {
	payload := new(model.CategoryPayload)

	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := payload.Validate(); err != nil {
		errs := utils.ValidationError(err.(validator.ValidationErrors))
		log.WithError(fiber.ErrBadRequest).Error("validate error :%w", errs)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errs,
		})
	}

	if err := h.service.Category.InsertCategory(ctx.Context(), payload.Name); err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
	})
}

func (h *CategoryHandler) UpdateCategory(ctx *fiber.Ctx) error {
	payload := new(model.CategoryPayload)
	name := ctx.Params("category_name")
	categoryName, err := url.PathUnescape(name)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := payload.Validate(); err != nil {
		errs := utils.ValidationError(err.(validator.ValidationErrors))
		log.WithError(fiber.ErrBadRequest).Error("validate error :%w", errs)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errs,
		})
	}

	resp, err := h.service.Category.UpdateCategory(ctx.Context(), payload.Name, categoryName)
	if err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    resp,
	})
}

func (h *CategoryHandler) DeleteCategory(ctx *fiber.Ctx) error {
	name := ctx.Params("category_name")
	categoryName, err := url.PathUnescape(name)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.Category.DeleteCategory(ctx.Context(), categoryName); err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

const (
	defaultLimit = 5
	maxLimit     = 100
	minOffset    = 0
)

func (h *CategoryHandler) GetCategories(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit", strconv.Itoa(defaultLimit)))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if limit <= 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	offset, err := strconv.Atoi(ctx.Query("offset", "0"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	params := model.PaginatinParams{
		Offset: offset,
		Limit:  limit,
	}

	resp, err := h.service.Category.GetCategories(ctx.Context(), params)
	if err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    resp,
	})
}
