package models

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/naumyegor/taxi-service/internal/db/interfaces"
	"github.com/pkg/errors"
)

const usersTableName = "users"

func (e *Env) CreateUser(user interfaces.User) error {
	queryString := fmt.Sprintf("INSERT INTO %s"+
		"(nickname, password, role_id)"+
		"VALUES ($1, $2, $3)", usersTableName)

	_, err := e.DB.Exec(queryString, user.Nickname, user.Password, user.RoleId)

	return err
}

func (e *Env) NicknameExists(nickname string) (bool, error) {
	var number int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE nickname = $1", usersTableName)

	err := e.DB.QueryRow(query, nickname).Scan(&number)
	if err != nil {
		return false, err
	}

	if number != 1 {
		return false, nil
	}

	return true, nil
}

func (e *Env) TokenExists(token string) (bool, error) {
	var number int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE token = $1", usersTableName)

	err := e.DB.QueryRow(query, token).Scan(&number)
	if err != nil {
		return false, err
	}

	if number != 1 {
		return false, nil
	}

	return true, nil
}

func (e *Env) GetUserByToken(token string) (interfaces.User, error) {
	var user interfaces.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE token = $1", usersTableName)

	err := e.DB.QueryRow(query, token).Scan(&user.ID, &user.Nickname, &user.Password,
		&user.Token, &user.RoleId)
	return user, err
}

func (e *Env) GetUserByNickname(nickname string) (interfaces.User, error) {
	var user interfaces.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE nickname = $1", usersTableName)

	err := e.DB.QueryRow(query, nickname).Scan(&user.ID, &user.Nickname, &user.Password,
		&user.Token, &user.RoleId)
	return user, err
}

func (e *Env) GetUserById(id int32) (interfaces.User, error) {
	var user interfaces.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTableName)

	err := e.DB.QueryRow(query, id).Scan(&user.ID, &user.Nickname, &user.Password,
		&user.Token, &user.RoleId)
	return user, err
}

func (e *Env) SeTokenByNickname(token, nickname string) error {
	query := fmt.Sprintf("UPDATE %s SET token = $1 WHERE nickname = $2", usersTableName)

	res, err := e.DB.Exec(query, token, nickname)
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

func (e *Env) UserHasCars(id int32) (bool, error) {
	var number int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE driver = $1", carsTableName)

	err := e.DB.QueryRow(query, id).Scan(&number)
	if err != nil {
		return false, err
	}

	if number != 1 {
		return false, nil
	}

	return true, nil
}
