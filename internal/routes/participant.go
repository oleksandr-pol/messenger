package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/messenger/internal/handlers"
	"github.com/oleksandr-pol/messenger/internal/models"
)

func Participant(r *mux.Router, db models.ParticipantRepository) {
	r.HandleFunc("/participant", handlers.NewParticipant(db)).Methods(http.MethodPost)
}
