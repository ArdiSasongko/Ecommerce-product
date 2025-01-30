package model

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type (
	ProductPayload struct {
		Name           string            `json:"name" validate:"required"`
		Description    string            `json:"description" validate:"required"`
		Price          float32           `json:"price" validate:"required,gte=5000"`
		Categories     []string          `json:"categories" validate:"required"`
		VariantProduct []VariantsPayload `json:"variants" validate:"dive"`
	}

	VariantsPayload struct {
		ProductID int32  `json:"product_id"`
		Color     string `json:"color" validate:"required"`
		Size      string `json:"size" validate:"required"`
		Quantity  int32  `json:"quantity" validate:"required,gte=1"`
	}
)

func (u ProductPayload) Validate() error {
	return Validate.Struct(u)
}

func (u VariantsPayload) Validate() error {
	return Validate.Struct(u)
}
