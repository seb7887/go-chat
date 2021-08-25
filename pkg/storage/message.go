package storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/challenge/pkg/models"
)

type MessageRepository interface {
	Create(msg models.NewMsgReq) (*models.Message, error)
	FindMessages(senderId uint, req models.GetMsgsReq) ([]QueryResult, error)
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
		Timestamp:   time.Now().UTC(),
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

type QueryResult struct {
	ID          uint
	SenderId    uint
	RecipientId uint
	ContentType string
	CreatedAt   time.Time
	Text        string
	ImageUrl    string
	Height      uint
	Width       uint
	VideoUrl    string
	Source      string
}

func (r *messageRepository) FindMessages(senderId uint, req models.GetMsgsReq) ([]QueryResult, error) {
	db := GetInstance()

	var result []QueryResult
	res := db.Raw(buildQuery(senderId, req.Recipient, req.Start, req.Limit)).Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	return result, nil
}

func buildQuery(senderId uint, recipientId uint, start uint, limit uint) string {
	selectClause := "SELECT m.*, t.text, i.url AS image_url, i.height, i.width, v.url AS video_url, v.source FROM messages m"
	textJoinClause := "LEFT OUTER JOIN texts t ON (t.message_id = m.id AND m.content_type = 'text')"
	imageJoinClause := "LEFT OUTER JOIN images i ON (i.message_id = m.id AND m.content_type = 'image')"
	videoJoinClause := "LEFT OUTER JOIN videos v ON (v.message_id = m.id AND m.content_type = 'video')"
	whereClause := fmt.Sprintf("WHERE m.id >= %d AND sender_id = %d AND recipient_id = %d", start, senderId, recipientId)
	orderByClause := fmt.Sprintf("ORDER BY m.id ASC LIMIT %d", limit)

	return fmt.Sprintf("%s %s %s %s %s %s;", selectClause, textJoinClause, imageJoinClause, videoJoinClause, whereClause, orderByClause)
}
