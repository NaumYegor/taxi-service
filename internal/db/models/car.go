package models

import (
	"fmt"
	"github.com/naumyegor/taxi-service/internal/db/interfaces"
	"github.com/pkg/errors"
)

const carsTableName = "cars"

func (e *Env) CreateCar(car interfaces.Car) error {
	queryString := fmt.Sprintf("INSERT INTO %s"+
		"(model, number, driver)"+
		"VALUES ($1, $2, $3)", carsTableName)

	_, err := e.DB.Exec(queryString, car.Model, car.Number, car.Driver)

	return err
}

func (e *Env) CarNumberExists(carNumber string) (bool, error) {
	var res int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE number = $1", carsTableName)

	err := e.DB.QueryRow(query, carNumber).Scan(&res)
	if err != nil {
		return false, err
	}

	if res != 1 {
		return false, nil
	}

	return true, nil
}

func (e *Env) CarIdExists(id int32) (bool, error) {
	var res int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id = $1", carsTableName)

	err := e.DB.QueryRow(query, id).Scan(&res)
	if err != nil {
		return false, err
	}

	if res != 1 {
		return false, nil
	}

	return true, nil
}

func (e *Env) SetDriverByCarId(CarId, DriverId int32) error {
	query := fmt.Sprintf("UPDATE %s SET driver = $1 WHERE id = $2", carsTableName)

	res, err := e.DB.Exec(query, DriverId, CarId)
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

func (e *Env) GetAvailableCars() ([]*interfaces.Car, error) {
	var cars []*interfaces.Car

	query := fmt.Sprintf("SELECT id, driver, model, number FROM cars " +
		"WHERE (cars.id NOT IN (SELECT car_id FROM orders WHERE completed=false AND orders.car_id!=NULL)) " +
		"OR (cars.id NOT IN (SELECT car_id FROM orders) AND driver!=NULL)")

	rows, err := e.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		car := new(interfaces.Car)
		if err := rows.Scan(&car.ID, &car.Driver, &car.Model, &car.Number); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}
