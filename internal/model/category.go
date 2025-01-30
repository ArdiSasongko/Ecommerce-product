package model

type (
	CategoryPayload struct {
		Name string `json:"name" validate:"required"`
	}
)

func (u CategoryPayload) Validate() error {
	return Validate.Struct(u)
}
