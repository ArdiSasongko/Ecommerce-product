package model

import "time"

type (
	CategoryPayload struct {
		Name string `json:"name" validate:"required"`
	}
)

func (u CategoryPayload) Validate() error {
	return Validate.Struct(u)
}

type (
	CategoryResponse struct {
		ID        int32     `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}
)
