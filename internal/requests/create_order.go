package requests

import "gopkg.in/go-playground/validator.v9"

type CreateOrderRequestAttributes struct {
	Info string `json:"info" validate:"required"`
}

func (co *CreateOrderRequestAttributes) Validate() error {
	v := validator.New()
	err := v.Struct(co)

	return err
}
