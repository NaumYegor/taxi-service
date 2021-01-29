package processors

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func NewSetDriverRequest(r *http.Request) (requests.SetDriverRequestAttributes, error) {
	var NewDriver requests.SetDriverRequestAttributes

	err := json.NewDecoder(r.Body).Decode(&NewDriver)
	if err != nil {
		return NewDriver, err
	}

	err = NewDriver.Validate()

	return NewDriver, err
}
