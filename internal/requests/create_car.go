package requests

import "gopkg.in/go-playground/validator.v9"

type CreateCarRequestAttributes struct {
	Model  string `json:"model" validate:"required"`
	Number string `json:"number" validate:"required"`
}

func (cr *CreateCarRequestAttributes) Validate() error {
	v := validator.New()
	err := v.Struct(cr)

	return err
}
