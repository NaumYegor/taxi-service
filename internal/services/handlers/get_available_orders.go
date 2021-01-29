package handlers

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func (he HandlerEnv) GetAvailableOrders(w http.ResponseWriter, _ *http.Request) {
	userModelEnv := models.Env(he)

	availableOrders, err := userModelEnv.GetAvailableOrders()
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	var responseList requests.GetAvailableOrdersResponseList
	responseList = make([]requests.GetAvailableOrdersResponseAttributes, 0)

	for _, orderPointer := range availableOrders {
		order := *orderPointer
		carAttributes := requests.GetAvailableOrdersResponseAttributes{
			ID:   order.ID,
			Info: order.Info,
		}
		responseList = append(responseList, carAttributes)
	}

	responseListJson, err := json.Marshal(responseList)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseListJson); err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
}
