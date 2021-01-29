package processors

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func NewCreateOrderRequest(r *http.Request) (requests.CreateOrderRequestAttributes, error) {
	var NewCar requests.CreateOrderRequestAttributes

	err := json.NewDecoder(r.Body).Decode(&NewCar)
	if err != nil {
		return NewCar, err
	}

	err = NewCar.Validate()

	return NewCar, err
}
