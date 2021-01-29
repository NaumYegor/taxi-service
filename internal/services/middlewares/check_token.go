package middlewares

import (
	"github.com/naumyegor/taxi-service/internal/config"
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/services/handlers"
	"net/http"
)

func CheckToken(next http.HandlerFunc, env config.Environment) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token == "" {
			handlers.ProcessBadRequest(&w, handlers.HandlerEnv(env), handlers.ErrorScheme{
				Type: handlers.TypeTokenError,
				Text: "you haven't add token to the headers",
			}, http.StatusUnauthorized)
			return
		}

		userModelEnv := models.Env(env)

		exists, err := userModelEnv.TokenExists(token)
		if err != nil {
			handlers.ProcessInternalError(&w, handlers.HandlerEnv(env), err)
			return
		}

		if !exists {
			handlers.ProcessBadRequest(&w, handlers.HandlerEnv(env), handlers.ErrorScheme{
				Type: handlers.TypeTokenError,
				Text: "this token does not exist",
			}, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
