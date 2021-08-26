package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
)

// CreateUser creates a new user
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.NewUserReq

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

	helpers.RespondJSON(w, models.NewUserResp{Id: user.ID})
}
