package main

import (
	"fmt"
	"net/http"

	"github.com/challenge/pkg/auth"
	"github.com/challenge/pkg/config"
	"github.com/challenge/pkg/controller"
	"github.com/challenge/pkg/services"
	"github.com/challenge/pkg/storage"
	log "github.com/sirupsen/logrus"
)

const (
	CheckEndpoint    = "/check"
	UsersEndpoint    = "/users"
	LoginEndpoint    = "/login"
	MessagesEndpoint = "/messages"
)

func main() {
	serverPort := config.GetConfig().ServerPort
	serverAddr := fmt.Sprintf(":%d", serverPort)
	userRepository := storage.NewUserRepository()
	userService := services.NewUserService(userRepository)

	// Perform DB migrations
	storage.Migrate()

	h := controller.NewHandler(userService)

	// Configure endpoints
	// Health
	http.HandleFunc(CheckEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		h.Check(w, r)
	})

	// Users
	http.HandleFunc(UsersEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		h.CreateUser(w, r)
	})

	// Auth
	http.HandleFunc(LoginEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		h.Login(w, r)
	})

	// Messages
	http.HandleFunc(MessagesEndpoint, auth.ValidateUser(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetMessages(w, r)
		case http.MethodPost:
			h.SendMessage(w, r)
		default:
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	}))

	// Start server
	log.Infof("Server started at port %d", serverPort)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
