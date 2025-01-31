package handler

import (
	"strconv"

	"github.com/ArdiSasongko/Ecommerce-product/internal/config/logger"
	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
	"github.com/ArdiSasongko/Ecommerce-product/internal/service"
	"github.com/ArdiSasongko/Ecommerce-product/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var log = logger.NewLogger()

type ProductHandler struct {
	service service.Service
}

func (h *ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	payload := new(model.ProductPayload)

	if err := ctx.BodyParser(payload); err != nil {
		log.WithError(fiber.ErrBadRequest).Error("body parse error :%w", err)
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

	if err := h.service.Product.CreateProduct(ctx.Context(), payload); err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
	})
}

func (h *ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	payload := new(model.ProductUpdatePayload)
	id := ctx.Params("productID")
	productID, err := strconv.Atoi(id)
	if err != nil {
		log.WithError(fiber.ErrBadRequest).Error("parsing error :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	payload.ProductID = int32(productID)

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

	resp, err := h.service.Product.UpdateProduct(ctx.Context(), payload)
	if err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
		"data":    resp,
	})
}

func (h *ProductHandler) UpdateVariant(ctx *fiber.Ctx) error {
	payload := new(model.VariantsUpdatePayload)
	id := ctx.Params("productID")
	Vid := ctx.Params("variantID")
	productID, err := strconv.Atoi(id)
	if err != nil {
		log.WithError(fiber.ErrBadRequest).Error("parsing error :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	variantID, err := strconv.Atoi(Vid)
	if err != nil {
		log.WithError(fiber.ErrBadRequest).Error("parsing error :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	payload.ProductID = int32(productID)
	payload.VariantID = int32(variantID)

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

	resp, err := h.service.Product.UpdateVariant(ctx.Context(), payload)
	if err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
		"data":    resp,
	})
}
