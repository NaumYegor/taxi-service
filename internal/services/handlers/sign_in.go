package handlers

import (
	crypto "crypto/rand"
	"encoding/hex"
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/services/processors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (he HandlerEnv) SignIn(w http.ResponseWriter, r *http.Request) {
	requestAttributes, err := processors.NewSignInRequest(r)
	if err != nil {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeRequestAttributesError,
		}, http.StatusBadRequest)
		return
	}
	userModelEnv := models.Env(he)

	exists, err := userModelEnv.NicknameExists(requestAttributes.Nickname)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if !exists {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: "this nickname doesn't exist",
		}, http.StatusNotFound)
		return
	}

	user, err := userModelEnv.GetUserByNickname(requestAttributes.Nickname)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(requestAttributes.Password))
	if err != nil {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: "wrong password",
		}, http.StatusUnauthorized)
		return
	}

	tokenBytes := make([]byte, 32)
	_, err = crypto.Read(tokenBytes)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	tokenString := hex.EncodeToString(tokenBytes)

	err = userModelEnv.SeTokenByNickname(tokenString, requestAttributes.Nickname)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	w.Header().Set("token", tokenString)

	w.WriteHeader(http.StatusOK)
}
