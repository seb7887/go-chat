package controller

import (
	"net/http"

	"github.com/challenge/pkg/services"
)

type Handler interface {
	Check(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	SendMessage(w http.ResponseWriter, r *http.Request)
	GetMessages(w http.ResponseWriter, r *http.Request)
}

// Handler provides the interface to handle different requests
type handler struct {
	userService services.UserService
}

func NewHandler(userService services.UserService) Handler {
	return &handler{
		userService: userService,
	}
}
