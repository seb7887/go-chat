package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/challenge/pkg/auth"
	"github.com/challenge/pkg/config"
	"github.com/challenge/pkg/controller"
	log "github.com/challenge/pkg/logger"
	"github.com/challenge/pkg/services"
	"github.com/challenge/pkg/storage"
)

const (
	CheckEndpoint    = "/check"
	UsersEndpoint    = "/users"
	LoginEndpoint    = "/login"
	MessagesEndpoint = "/messages"
)

func main() {
	var (
		serverPort        = config.GetConfig().ServerPort
		serverAddr        = fmt.Sprintf(":%d", serverPort)
		userRepository    = storage.NewUserRepository()
		messageRepository = storage.NewMessageRepository()
		userService       = services.NewUserService(userRepository)
		messageService    = services.NewMessageService(messageRepository)
		h                 = controller.NewHandler(userService, messageService)
	)

	// Setup logger
	log.Setup()

	// Mux Router
	mux := http.NewServeMux()

	// Configure endpoints
	// Health
	mux.HandleFunc(CheckEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		h.Check(w, r)
	})

	// Users
	mux.HandleFunc(UsersEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		h.CreateUser(w, r)
	})

	// Auth
	mux.HandleFunc(LoginEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		h.Login(w, r)
	})

	// Messages
	mux.HandleFunc(MessagesEndpoint, auth.ValidateUser(func(w http.ResponseWriter, r *http.Request) {
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

	// Create HTTP server
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Infof("Server started at port %d", serverPort)
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err.Error())
			}
		}
	}()

	// Gracefully shutdown
	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown")
	}

	log.Info("Server exiting")
}
