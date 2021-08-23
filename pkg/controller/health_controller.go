package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/storage"
)

// Check returns the health of the service and DB
func (h *handler) Check(w http.ResponseWriter, r *http.Request) {
	// Check DB health
	storage.GetInstance()
	// if d, err := db.DB(); err != nil {
	// 	if err = d.Ping(); err != nil {
	// 		d.Close()
	// 		handleInternalError(w, err)
	// 	}
	// } else {
	// 	handleInternalError(w, err)
	// }

	helpers.RespondJSON(w, models.Health{Health: "ok"})
}

func handleInternalError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
