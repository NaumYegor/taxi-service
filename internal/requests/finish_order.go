package requests

import "gopkg.in/go-playground/validator.v9"

type FinishOrderRequestAttributes struct {
	OrderId int `json:"order_id" validate:"required"`
}

func (dr *FinishOrderRequestAttributes) Validate() error {
	v := validator.New()
	err := v.Struct(dr)

	return err
}
