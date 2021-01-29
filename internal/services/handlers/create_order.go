package handlers

import (
	"github.com/naumyegor/taxi-service/internal/db/interfaces"
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/requests"
	"github.com/naumyegor/taxi-service/internal/services/processors"
	"net/http"
)

func (he HandlerEnv) CreateOrder(w http.ResponseWriter, r *http.Request) {
	request, err := processors.NewCreateOrderRequest(r)
	if err != nil {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeRequestAttributesError,
		}, http.StatusBadRequest)
		return
	}
	token := r.Header.Get("token")
	dbEnv := models.Env(he)

	client, err := dbEnv.GetUserByToken(token)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	if !requests.IsClient(int(client.RoleId)) {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypePermissionsError,
			Text: "you are not able to create orders. you are not a client",
		}, http.StatusForbidden)
		return
	}

	clientHasUncompletedTrip, err := dbEnv.ClientHasUncompletedTrip(client.ID)
	if err != nil {
		ProcessInternalError(&w, he, err)
	}
	if clientHasUncompletedTrip {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypePermissionsError,
			Text: "you are not able to create orders cause of uncompleted trips",
		}, http.StatusForbidden)
		return
	}

	newOrder := interfaces.Order{
		ClientId: client.ID,
		Info:     request.Info,
	}

	err = dbEnv.CreateOrder(newOrder)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
