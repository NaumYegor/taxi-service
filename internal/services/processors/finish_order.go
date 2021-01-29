package processors

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func NewFinishOrderRequest(r *http.Request) (requests.FinishOrderRequestAttributes, error) {
	var finishing requests.FinishOrderRequestAttributes

	err := json.NewDecoder(r.Body).Decode(&finishing)
	if err != nil {
		return finishing, err
	}

	err = finishing.Validate()

	return finishing, err
}
