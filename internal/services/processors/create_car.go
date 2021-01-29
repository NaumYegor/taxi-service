package processors

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func NewCreateCarRequest(r *http.Request) (requests.CreateCarRequestAttributes, error) {
	var NewCar requests.CreateCarRequestAttributes

	err := json.NewDecoder(r.Body).Decode(&NewCar)
	if err != nil {
		return NewCar, err
	}

	err = NewCar.Validate()

	return NewCar, err
}
