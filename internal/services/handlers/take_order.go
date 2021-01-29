package handlers

import (
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/services/processors"
	"net/http"
)

func (he HandlerEnv) TakeOrder(w http.ResponseWriter, r *http.Request) {
	request, err := processors.NewTakeOrderRequest(r)
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

	userHasCars, err := dbEnv.UserHasCars(user.ID)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if !userHasCars {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeRequestAttributesError,
			Text: "you don't have cars",
		}, http.StatusBadRequest)
		return
	}

	orderExists, err := dbEnv.OrderExists(int32(request.OrderId))
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if !orderExists {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeRequestAttributesError,
			Text: "this order doesn't exist",
		}, http.StatusBadRequest)
		return
	}

	orderIsAvailable, err := dbEnv.OrderIsAvailable(int32(request.OrderId))
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if !orderIsAvailable {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeRequestAttributesError,
			Text: "this order isn't available",
		}, http.StatusBadRequest)
		return
	}

	if err = dbEnv.TakeOrder(int32(request.OrderId), user.ID); err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
