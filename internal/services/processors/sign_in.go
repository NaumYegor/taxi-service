package processors

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/requests"
	"net/http"
)

func NewSignInRequest(r *http.Request) (requests.SignInRequestAttributes, error) {
	var User requests.SignInRequestAttributes

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		return User, err
	}

	err = User.Validate()

	return User, err
}
