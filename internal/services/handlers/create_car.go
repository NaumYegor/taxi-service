package handlers

import (
	"github.com/naumyegor/taxi-service/internal/db/interfaces"
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/requests"
	"github.com/naumyegor/taxi-service/internal/services/processors"
	"net/http"
)

func (he HandlerEnv) CreateCar(w http.ResponseWriter, r *http.Request) {
	request, err := processors.NewCreateCarRequest(r)
	if err != nil {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeRequestAttributesError,
		}, http.StatusBadRequest)
		return
	}
	token := r.Header.Get("token")
	dbEnv := models.Env(he)

	user, err := dbEnv.GetUserByToken(token)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	if !requests.IsAdmin(int(user.RoleId)) {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypePermissionsError,
			Text: "you are not able to create cars",
		}, http.StatusForbidden)
		return
	}

	carNumberExists, err := dbEnv.CarNumberExists(request.Number)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if carNumberExists {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: "car with this number already exists",
		}, http.StatusConflict)
		return
	}

	newCar := interfaces.Car{
		Number: request.Number,
		Model:  request.Model,
		Driver: nil,
	}
	err = dbEnv.CreateCar(newCar)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
