package interfaces

type Car struct {
	ID     int32  `db:"id"`
	Driver *int32 `db:"driver"`
	Model  string `db:"model"`
	Number string `db:"number"`
}

type Cars interface {
	CreateCar(car Car) error
	CarNumberExists(carNumber string) (bool, error)
	CarIdExists(id int32) (bool, error)
	SetDriverByCarId(CarId, DriverId int32) error
	GetAvailableCars() ([]Car, error)
}
