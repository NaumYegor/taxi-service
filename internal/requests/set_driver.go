package requests

import "gopkg.in/go-playground/validator.v9"

type SetDriverRequestAttributes struct {
	CarId    int `json:"car_id" validate:"required"`
	DriverId int `json:"driver_id" validate:"required"`
}

func (dr *SetDriverRequestAttributes) Validate() error {
	v := validator.New()
	err := v.Struct(dr)

	return err
}
