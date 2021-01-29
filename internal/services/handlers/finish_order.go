package handlers

import (
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/services/processors"
	"net/http"
)

func (he HandlerEnv) FinishOrder(w http.ResponseWriter, r *http.Request) {
	request, err := processors.NewFinishOrderRequest(r)
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

	driverOnOrder, err := dbEnv.DriverOnOrder(user.ID, int32(request.OrderId))
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if !driverOnOrder {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeRequestAttributesError,
			Text: "you are not on this order",
		}, http.StatusBadRequest)
		return
	}

	if err = dbEnv.FinishOrder(int32(request.OrderId)); err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
