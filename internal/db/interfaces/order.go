package interfaces

type Order struct {
	ID        int32  `db:"id"`
	ClientId  int32  `db:"client_id"`
	CarId     *int32 `db:"car_id"`
	Info      string `db:"info"`
	Completed bool   `db:"completed"`
}

type Orders interface {
	CreateOrder(order Order) error
	ClientHasUncompletedTrip(clientId int32) (bool, error)
	GetAvailableOrders() ([]*Order, error)
}
