package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
)

// SendMessage send a message from one user to another
func (h *handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	// TODO: Send a New Message
	helpers.RespondJSON(w, models.Message{})
}

// GetMessages get the messages from the logged user to a recipient
func (h *handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Retrieve list of Messages
	helpers.RespondJSON(w, []*models.Message{{}})
}
