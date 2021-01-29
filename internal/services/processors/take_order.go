package processors

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func NewTakeOrderRequest(r *http.Request) (requests.TakeOrderRequestAttributes, error) {
	var taking requests.TakeOrderRequestAttributes

	err := json.NewDecoder(r.Body).Decode(&taking)
	if err != nil {
		return taking, err
	}

	err = taking.Validate()

	return taking, err
}
