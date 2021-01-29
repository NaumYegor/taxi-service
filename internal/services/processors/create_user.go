package processors

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func NewCreateUserRequest(r *http.Request) (requests.CreateUserRequestAttributes, error) {
	var NewUser requests.CreateUserRequestAttributes

	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		return NewUser, err
	}

	err = NewUser.Validate()

	return NewUser, err
}
