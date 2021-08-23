package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
)

type loginReq struct {
	Username string
	Password string
}

// Login authenticates a user and returns a token
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials loginReq

	helpers.UnmarshallBody(r, &credentials)

	if err := helpers.ValidateRequestBody(credentials); err != nil {
		helpers.HandleError(w, "Invalid request body")
		return
	}

	// Authenticate user
	authData, err := h.userService.LoginUser(credentials.Username, credentials.Password)
	if err != nil {
		helpers.HandleError(w, err.Error())
		return
	}

	helpers.RespondJSON(w, authData)
}
