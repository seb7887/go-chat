package main

import (
	"github.com/challenge/pkg/config"
	"github.com/challenge/pkg/controller"
	log "github.com/challenge/pkg/logger"
	"github.com/challenge/pkg/server"
	"github.com/challenge/pkg/services"
	"github.com/challenge/pkg/storage"
)

func main() {
	var (
		serverPort        = config.GetConfig().ServerPort
		userRepository    = storage.NewUserRepository()
		messageRepository = storage.NewMessageRepository()
		userService       = services.NewUserService(userRepository)
		messageService    = services.NewMessageService(messageRepository)
		h                 = controller.NewHandler(userService, messageService)
	)

	// Setup logger
	log.Setup()

	server.Serve(serverPort, h)
}
