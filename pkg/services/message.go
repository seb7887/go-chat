package services

import (
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/storage"
)

type MessageService interface {
	AddMessage(msg models.NewMsgReq) (*models.Message, error)
}

type messageService struct {
	repository storage.MessageRepository
}

func NewMessageService(repository storage.MessageRepository) MessageService {
	return &messageService{
		repository,
	}
}

func (s *messageService) AddMessage(msg models.NewMsgReq) (*models.Message, error) {
	// TODO: verify content type
	// TODO: verify content source type
	newMsg, err := s.repository.Create(msg)
	if err != nil {
		return nil, err
	}

	return newMsg, nil
}
