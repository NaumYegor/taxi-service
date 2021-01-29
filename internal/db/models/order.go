package models

import (
	"fmt"
	"github.com/naumyegor/taxi-service/internal/db/interfaces"
	"github.com/pkg/errors"
)

const ordersTableName = "orders"

func (e *Env) CreateOrder(order interfaces.Order) error {
	queryString := fmt.Sprintf("INSERT INTO %s"+
		"(client_id, info, car_id) VALUES ($1, $2, NULL)", ordersTableName)

	_, err := e.DB.Exec(queryString, order.ClientId, order.Info)

	return err
}

func (e *Env) ClientHasUncompletedTrip(clientId int32) (bool, error) {
	var number int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s "+
		"WHERE completed = false AND client_id = $1", ordersTableName)

	err := e.DB.QueryRow(query, clientId).Scan(&number)
	if err != nil {
		return false, err
	}

	if number != 1 {
		return false, nil
	}

	return true, nil
}

func (e *Env) GetAvailableOrders() ([]*interfaces.Order, error) {
	var orders []*interfaces.Order

	query := fmt.Sprintf("SELECT id, info FROM %s "+
		"WHERE car_id IS NULL AND completed=false", ordersTableName)

	rows, err := e.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		order := new(interfaces.Order)
		if err := rows.Scan(&order.ID, &order.Info); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (e *Env) OrderExists(id int32) (bool, error) {
	var number int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s "+
		"WHERE id = $1", ordersTableName)

	err := e.DB.QueryRow(query, id).Scan(&number)
	if err != nil {
		return false, err
	}

	if number != 1 {
		return false, nil
	}

	return true, nil
}

func (e *Env) OrderIsAvailable(id int32) (bool, error) {
	var number int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s "+
		"WHERE id = $1 AND car_id IS NULL AND completed=false", ordersTableName)

	err := e.DB.QueryRow(query, id).Scan(&number)
	if err != nil {
		return false, err
	}

	if number != 1 {
		return false, nil
	}

	return true, nil
}

func (e *Env) TakeOrder(orderId, driverId int32) error {
	query := fmt.Sprintf("UPDATE %s SET car_id = (SELECT id FROM cars WHERE driver=$1) "+
		"WHERE id=$2", ordersTableName)

	res, err := e.DB.Exec(query, driverId, orderId)
	if err != nil {
		return errors.Wrap(err, "unable to update row")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "unable to get affected rows")
	}
	if rowsAffected == 0 {
		return errors.New("Raw somehow hasn't been updated")
	}
	return nil
}

func (e *Env) FinishOrder(id int32) error {
	query := fmt.Sprintf("UPDATE %s SET completed=true WHERE id=$1", ordersTableName)

	res, err := e.DB.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "unable to update row")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "unable to get affected rows")
	}
	if rowsAffected == 0 {
		return errors.New("Raw somehow hasn't been updated")
	}
	return nil
}

func (e *Env) DriverOnOrder(driverId, orderId int32) (bool, error) {
	var number int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s "+
		"WHERE id = $1 AND car_id IN (SELECT id FROM cars WHERE driver=$2)", ordersTableName)

	err := e.DB.QueryRow(query, orderId, driverId).Scan(&number)
	if err != nil {
		return false, err
	}

	if number != 1 {
		return false, nil
	}

	return true, nil
}
