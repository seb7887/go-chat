package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
)

// Check returns the health of the service and DB
func (h *handler) Check(w http.ResponseWriter, r *http.Request) {
	helpers.RespondJSON(w, models.Health{Health: "ok"})
}

func handleInternalError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
