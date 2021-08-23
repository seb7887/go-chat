package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
)

type newUserReq struct {
	Username string
	Password string
}

type newUserResp struct {
	Id uint `json:"id"`
}

// CreateUser creates a new user
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser newUserReq

	helpers.UnmarshallBody(r, &newUser)

	if err := helpers.ValidateRequestBody(newUser); err != nil {
		helpers.HandleError(w, "Invalid request body")
		return
	}

	// Create a new user
	user, err := h.userService.AddUser(newUser.Username, newUser.Password)
	if err != nil {
		helpers.HandleError(w, "Error creating new user")
		return
	}

	helpers.RespondJSON(w, newUserResp{Id: user.ID})
}
