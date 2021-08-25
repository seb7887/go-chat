package services

import (
	"fmt"

	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/storage"
)

type MessageService interface {
	AddMessage(msg models.NewMsgReq) (*models.Message, error)
	GetMessages(senderId uint, req models.GetMsgsReq) ([]*models.MessageResp, error)
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
	newMsg, err := s.repository.Create(msg)
	if err != nil {
		return nil, err
	}

	return newMsg, nil
}

func (s *messageService) GetMessages(senderId uint, req models.GetMsgsReq) ([]*models.MessageResp, error) {
	queryResult, err := s.repository.FindMessages(senderId, req)
	if err != nil {
		return nil, err
	}

	response, err := formatMessagesResponse(queryResult)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func formatMessagesResponse(result []storage.QueryResult) ([]*models.MessageResp, error) {
	var resp []*models.MessageResp

	for _, v := range result {
		r, err := formatMessageResp(v)
		if err != nil {
			return resp, err
		}
		resp = append(resp, &r)
	}

	return resp, nil
}

func formatMessageResp(result storage.QueryResult) (models.MessageResp, error) {
	var resp models.MessageResp
	resp.ID = result.ID
	resp.Sender = result.SenderId
	resp.Recipient = result.RecipientId
	resp.Timestamp = result.CreatedAt
	resp.Content.Type = result.ContentType

	switch resp.Content.Type {
	case "text":
		resp.Content.Text = result.Text
	case "image":
		resp.Content.Url = result.ImageUrl
		resp.Content.Height = result.Height
		resp.Content.Width = result.Width
	case "video":
		resp.Content.Url = result.VideoUrl
		resp.Content.Source = result.Source
	default:
		return resp, fmt.Errorf("Error formating message response")
	}

	return resp, nil
}
