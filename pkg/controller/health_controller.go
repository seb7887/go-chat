package controller

import (
	"net/http"
	"os"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
)

// Check returns the health of the service and DB
func (h *handler) Check(w http.ResponseWriter, r *http.Request) {
	// Check if database is alive
	if _, err := os.Stat("./chat.db"); os.IsNotExist(err) {
		handleInternalError(w, err)
		return
	}
	helpers.RespondJSON(w, models.Health{Health: "ok"})
}

func handleInternalError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
