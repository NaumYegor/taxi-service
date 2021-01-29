package services

import (
	"github.com/gorilla/mux"
	"github.com/naumyegor/taxi-service/internal/config"
	"github.com/naumyegor/taxi-service/internal/db"
	"github.com/naumyegor/taxi-service/internal/services/handlers"
	"github.com/naumyegor/taxi-service/internal/services/middlewares"
	"net/http"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	env := config.NewEnvironment()
	var handlerEnv Handlers = handlers.HandlerEnv(*env)

	//TODO: Provide a more convenient way to use migrations
	db.Migrate(env)
	handlers.CreateRootUser(handlers.HandlerEnv(*env))

	//TODO: Rename some routes according to the REST conventions
	r.HandleFunc("/users",
		middlewares.CheckToken(handlerEnv.CreateUser, *env)).Methods(http.MethodPost)
	r.HandleFunc("/authentication", handlerEnv.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/cars", middlewares.CheckToken(handlerEnv.CreateCar, *env)).Methods(http.MethodPost)
	r.HandleFunc("/cars", middlewares.CheckToken(handlerEnv.SetDriver, *env)).Methods(http.MethodPut)
	r.HandleFunc("/cars",
		middlewares.CheckToken(handlerEnv.GetAvailableCars, *env)).Methods(http.MethodGet)
	r.HandleFunc("/orders", middlewares.CheckToken(handlerEnv.CreateOrder, *env)).Methods(http.MethodPost)
	r.HandleFunc("/orders", middlewares.CheckToken(handlerEnv.GetAvailableOrders, *env)).Methods(http.MethodGet)
	r.HandleFunc("/orders", middlewares.CheckToken(handlerEnv.TakeOrder, *env)).Methods(http.MethodPut)
	r.HandleFunc("/orders", middlewares.CheckToken(handlerEnv.FinishOrder, *env)).Methods(http.MethodDelete)

	return r
}
