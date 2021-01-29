package requests

import "gopkg.in/go-playground/validator.v9"

type TakeOrderRequestAttributes struct {
	OrderId int `json:"order_id" validate:"required"`
}

func (dr *TakeOrderRequestAttributes) Validate() error {
	v := validator.New()
	err := v.Struct(dr)

	return err
}
