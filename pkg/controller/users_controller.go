package controller

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/helpers"
)

type newUserReq struct {
	Username string
	Password string
}

type newUserResp struct {
	Id int `json:"id"`
}

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var newUser newUserReq
	err := decoder.Decode(&newUser)
	if err != nil {
		helpers.HandleError(w, "Invalid request body")
		return
	}

	// TODO: Create a New User
	helpers.RespondJSON(w, newUserResp{Id: 0})
}
