package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
)

var (
	validContentTypes = map[string]string{
		"text":  "text",
		"image": "image",
		"video": "video",
	}
	validSourceTypes = map[string]string{
		"youtube": "youtube",
		"vimeo":   "vimeo",
	}
)

type newMsgResp struct {
	Id        uint      `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

// SendMessage send a message from one user to another
func (h *handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(float64)
	var newMsg models.NewMsgReq

	helpers.UnmarshallBody(r, &newMsg)

	// Validate request body
	err := helpers.ValidateRequestBody(newMsg)
	if err != nil {
		helpers.HandleError(w, "Invalid request body")
		return
	}

	// SenderID must be the same as UserID
	if newMsg.Sender != uint(userId) {
		helpers.HandleError(w, "Sender must be the user who is logged in")
		return
	}

	// Validate content type and content source (if present)
	err = validateMsgContent(newMsg.Content)
	if err != nil {
		helpers.HandleError(w, err.Error())
		return
	}

	message, err := h.messageService.AddMessage(newMsg)
	if err != nil {
		helpers.HandleError(w, "Error creating new message")
		return
	}

	helpers.RespondJSON(w, newMsgResp{Id: message.ID, Timestamp: message.Timestamp})
}

func validateMsgContent(content models.MsgContent) error {
	contentType := strings.ToLower(content.Type)
	if _, exists := validContentTypes[contentType]; !exists {
		return fmt.Errorf("Invalid content type")
	}

	if contentType == "video" {
		if _, exists := validSourceTypes[strings.ToLower(content.Source)]; !exists {
			return fmt.Errorf("Invalid content source")
		}
	}

	return nil
}

// GetMessages get the messages from the logged user to a recipient
func (h *handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(float64)

	// Parse query string params
	req, err := parseQueryStrings(r)
	if err != nil {
		helpers.HandleError(w, "Error parsing query string params")
		return
	}

	if err := helpers.ValidateRequestBody(req); err != nil {
		helpers.HandleError(w, "Invalid request body")
		return
	}

	// Retrieve list of Messages
	messages, err := h.messageService.GetMessages(uint(userId), *req)
	if err != nil {
		helpers.HandleError(w, err.Error())
		return
	}

	helpers.RespondJSON(w, messages)
}

func parseQueryStrings(r *http.Request) (*models.GetMsgsReq, error) {
	query := r.URL.Query()

	recipient, err := strconv.ParseUint(query["recipient"][0], 10, 32)
	start, err := strconv.ParseUint(query["start"][0], 10, 32)
	limit, err := strconv.ParseUint(query["limit"][0], 10, 32)

	if err != nil {
		return nil, err
	}

	return &models.GetMsgsReq{
		Recipient: uint(recipient),
		Start:     uint(start),
		Limit:     uint(limit),
	}, nil
}
