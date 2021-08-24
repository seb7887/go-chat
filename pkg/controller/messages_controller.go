package controller

import (
	"net/http"
	"time"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
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

	if err := helpers.ValidateRequestBody(newMsg); err != nil {
		helpers.HandleError(w, "Invalid request body")
		return
	}

	// SenderID must be the same as UserID
	if newMsg.Sender != uint(userId) {
		helpers.HandleError(w, "Sender must be the user who is logged in")
		return
	}

	message, err := h.messageService.AddMessage(newMsg)
	if err != nil {
		helpers.HandleError(w, "Error creating new message")
		return
	}

	helpers.RespondJSON(w, newMsgResp{Id: message.ID, Timestamp: message.Timestamp})
}

// GetMessages get the messages from the logged user to a recipient
func (h *handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Retrieve list of Messages
	helpers.RespondJSON(w, []*models.Message{{}})
}
