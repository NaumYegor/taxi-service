package requests

import "gopkg.in/go-playground/validator.v9"

type CreateUserRequestAttributes struct {
	Role     string `json:"role" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (nur *CreateUserRequestAttributes) Validate() error {
	v := validator.New()
	err := v.Struct(nur)

	return err
}
