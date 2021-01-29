package handlers

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func (he HandlerEnv) GetAvailableCars(w http.ResponseWriter, _ *http.Request) {
	userModelEnv := models.Env(he)

	availableCars, err := userModelEnv.GetAvailableCars()
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	var responseList requests.GetAvailableCarsResponseList
	responseList = make([]requests.GetAvailableCarsResponseAttributes, 0)

	for _, carPointer := range availableCars {
		car := *carPointer
		carAttributes := requests.GetAvailableCarsResponseAttributes{
			ID:     car.ID,
			Driver: *car.Driver,
			Model:  car.Model,
			Number: car.Number,
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
