package handlers

import (
	"encoding/json"
	"github.com/naumyegor/taxi-service/internal/config"
	"github.com/naumyegor/taxi-service/internal/db/interfaces"
	"github.com/naumyegor/taxi-service/internal/db/models"
	"github.com/naumyegor/taxi-service/internal/requests"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	TypeValidationError        = "validation"
	TypeRequestAttributesError = "request_attributes"
	TypePermissionsError       = "permissions_error"
	TypeTokenError             = "token"
	rootUserNickname           = "root"
	rootUserPassword           = "root"
)

type HandlerEnv config.Environment

type ErrorScheme struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

func ProcessInternalError(w *http.ResponseWriter, env HandlerEnv, err error) {
	http.Error(*w, "", http.StatusInternalServerError)
	env.Logger.Println(err)
}

func ProcessBadRequest(w *http.ResponseWriter, env HandlerEnv, errScheme ErrorScheme, StatusCode int) {
	errBytes, err := json.Marshal(errScheme)
	if err != nil {
		ProcessInternalError(w, env, err)
	}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Write(errBytes)

	env.Logger.Println(string(errBytes))
}

func CreateRootUser(he HandlerEnv) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rootUserPassword), bcrypt.DefaultCost)
	if err != nil {
		he.Logger.Println("Failed to create root user: bcrypt")
		return
	}

	rootRole, err := requests.GetRoleId("Admin")
	if err != nil {
		he.Logger.Println("Failed to create root user: get role")
		return
	}

	rootUser := interfaces.User{
		Nickname: rootUserNickname,
		Password: hashedPassword,
		RoleId:   int32(rootRole),
	}

	dbModel := models.Env(he)
	err = dbModel.CreateUser(rootUser)
	if err != nil {
		he.Logger.Println("Failed to create root user: db's query")
		return
	}

	he.Logger.Println("Created root user successfully.")
}
