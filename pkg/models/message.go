package models

import (
	"time"
)

type Message struct {
	ID          uint
	SenderID    uint      `gorm:"column:sender_id"`
	RecipientID uint      `gorm:"recipient_id"`
	ContentType string    `gorm:"column:content_type"`
	Timestamp   time.Time `gorm:"column:created_at"`
}

type Text struct {
	ID        uint
	MessageID uint `gorm:"column:message_id"`
	Text      string
}

type Image struct {
	ID        uint
	MessageID uint `gorm:"column:message_id"`
	Url       string
	Height    uint
	Width     uint
}

type Video struct {
	ID        uint
	MessageID uint `gorm:"column:message_id"`
	Url       string
	Source    string
}

type MsgContent struct {
	Type   string `json:"type"`
	Text   string `json:"text,omitempty"`
	Url    string `json:"url,omitempty"`
	Height uint   `json:"height,omitempty"`
	Width  uint   `json:"width,omitempty"`
	Source string `json:"source,omitempty"`
}

type NewMsgReq struct {
	Sender    uint
	Recipient uint
	Content   MsgContent
}

type GetMsgsReq struct {
	Recipient uint
	Start     uint
	Limit     uint
}

type MessageResp struct {
	ID        uint       `json:"id"`
	Timestamp time.Time  `json:"timestamp"`
	Sender    uint       `json:"sender"`
	Recipient uint       `json:"recipient"`
	Content   MsgContent `json:"content"`
}
