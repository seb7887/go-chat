package storage

import (
	"strings"
	"time"

	"github.com/challenge/pkg/models"
)

type MessageRepository interface {
	Create(msg models.NewMsgReq) (*models.Message, error)
}

type messageRepository struct{}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (r *messageRepository) Create(msg models.NewMsgReq) (*models.Message, error) {
	message := models.Message{
		SenderID:    msg.Sender,
		RecipientID: msg.Recipient,
		ContentType: msg.Content.Type,
		Timestamp:   time.Now(),
	}

	db := GetInstance()
	res := db.Create(&message)

	if res.Error != nil {
		return nil, res.Error
	}

	contentType := strings.ToLower(msg.Content.Type)

	// Create content based on contentType
	switch contentType {
	case "text":
		textContent := models.Text{
			MessageID: message.ID,
			Text:      msg.Content.Text,
		}

		res = db.Create(&textContent)
	case "image":
		imgContent := models.Image{
			MessageID: message.ID,
			Url:       msg.Content.Url,
			Height:    msg.Content.Height,
			Width:     msg.Content.Width,
		}

		res = db.Create(&imgContent)
	case "video":
		videoContent := models.Video{
			MessageID: message.ID,
			Url:       msg.Content.Url,
			Source:    msg.Content.Source,
		}

		res = db.Create(&videoContent)
	}

	if res.Error != nil {
		db.Delete(&message)
		return nil, res.Error
	}

	return &message, nil
}
