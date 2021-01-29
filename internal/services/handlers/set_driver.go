package handlers

import (
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/requests"
	"github.com/naumyegor/taxi-service/internal/services/processors"
	"net/http"
)

func (he HandlerEnv) SetDriver(w http.ResponseWriter, r *http.Request) {
	request, err := processors.NewSetDriverRequest(r)
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
			Text: "you are not able to set drivers for cars",
		}, http.StatusForbidden)
		return
	}

	carIdExists, err := dbEnv.CarIdExists(int32(request.CarId))
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if !carIdExists {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: "car with this id doesn't exists",
		}, http.StatusNotFound)
		return
	}

	driver, err := dbEnv.GetUserById(int32(request.DriverId))
	if err != nil {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: "user with this id doesn't exists",
		}, http.StatusNotFound)
		return
	}
	if !requests.IsDriver(int(driver.RoleId)) {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: "user with this id isn't a driver",
		}, http.StatusConflict)
		return
	}

	driveHasCars, err := dbEnv.UserHasCars(driver.ID)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if driveHasCars {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: "driver with this id has a car",
		}, http.StatusConflict)
		return
	}

	err = dbEnv.SetDriverByCarId(int32(request.CarId), int32(request.DriverId))
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
