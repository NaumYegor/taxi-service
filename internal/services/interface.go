package services

import (
	"net/http"
)

type Handlers interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	CreateCar(w http.ResponseWriter, r *http.Request)
	SetDriver(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
	GetAvailableCars(w http.ResponseWriter, r *http.Request)
	GetAvailableOrders(w http.ResponseWriter, r *http.Request)
	TakeOrder(w http.ResponseWriter, r *http.Request)
	FinishOrder(w http.ResponseWriter, r *http.Request)
}
