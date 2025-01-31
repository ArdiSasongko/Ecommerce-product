package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type (
	ProductPayload struct {
		Name           string            `json:"name" validate:"required"`
		Description    string            `json:"description" validate:"required"`
		Price          float32           `json:"price" validate:"required,gte=5000,numeric"`
		Categories     []string          `json:"categories" validate:"required"`
		VariantProduct []VariantsPayload `json:"variants" validate:"dive"`
	}

	VariantsPayload struct {
		ProductID int32  `json:"product_id"`
		Color     string `json:"color" validate:"required"`
		Size      string `json:"size" validate:"required"`
		Quantity  int32  `json:"quantity" validate:"required,gte=1"`
	}

	ProductUpdatePayload struct {
		ProductID   int32    `json:"product_id"`
		Name        *string  `json:"name" validate:"omitempty,min=1"`
		Description *string  `json:"description" validate:"omitempty,min=1"`
		Price       *float32 `json:"price" validate:"omitempty,gte=5000,numeric"`
	}

	VariantsUpdatePayload struct {
		ProductID int32   `json:"product_id"`
		VariantID int32   `json:"variant_id"`
		Color     *string `json:"color" validate:"omitempty"`
		Size      *string `json:"size" validate:"omitempty"`
		Quantity  *int32  `json:"quantity" validate:"omitempty,gte=1"`
	}
)

func (u ProductPayload) Validate() error {
	return Validate.Struct(u)
}

func (u VariantsPayload) Validate() error {
	return Validate.Struct(u)
}

func (u ProductUpdatePayload) Validate() error {
	return Validate.Struct(u)
}

func (u VariantsUpdatePayload) Validate() error {
	return Validate.Struct(u)
}

type (
	ProductUpdateResponse struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       float32   `json:"price"`
		UpdateAt    time.Time `json:"update_at"`
	}

	VariantUpdateResponse struct {
		Color     string    `json:"color"`
		Size      string    `json:"size"`
		Quantity  int32     `json:"quantity"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
