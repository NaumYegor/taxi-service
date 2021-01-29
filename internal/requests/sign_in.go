package requests

import "gopkg.in/go-playground/validator.v9"

type SignInRequestAttributes struct {
	Nickname string `json:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (sir *SignInRequestAttributes) Validate() error {
	v := validator.New()
	err := v.Struct(sir)

	return err
}
