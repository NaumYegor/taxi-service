package handlers

import (
	"github.com/naumyegor/taxi-service/internal/db/interfaces"
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/requests"
	"github.com/naumyegor/taxi-service/internal/services/processors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (he HandlerEnv) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestAttributes, err := processors.NewCreateUserRequest(r)
	if err != nil {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeRequestAttributesError,
		}, http.StatusBadRequest)
		return
	}
	token := r.Header.Get("token")
	userModelEnv := models.Env(he)

	user, err := userModelEnv.GetUserByToken(token)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	if !requests.IsAdmin(int(user.RoleId)) {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypePermissionsError,
			Text: "you are not able to create users",
		}, http.StatusForbidden)
		return
	}

	nicknameExists, err := userModelEnv.NicknameExists(requestAttributes.Nickname)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}
	if nicknameExists {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: "this nickname already exists",
		}, http.StatusConflict)
		return
	}

	newUserRoleId, err := requests.GetRoleId(requestAttributes.Role)
	if err != nil {
		ProcessBadRequest(&w, he, ErrorScheme{
			Type: TypeValidationError,
			Text: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestAttributes.Password), bcrypt.DefaultCost)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	newUser := interfaces.User{
		Nickname: requestAttributes.Nickname,
		Password: hashedPassword,
		RoleId:   int32(newUserRoleId),
	}

	err = userModelEnv.CreateUser(newUser)
	if err != nil {
		ProcessInternalError(&w, he, err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	return
}
