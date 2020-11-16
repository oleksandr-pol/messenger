package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/messenger/internal/handlers"
	"github.com/oleksandr-pol/messenger/internal/models"
)

func Message(r *mux.Router, db models.MessageRepository) {
	r.HandleFunc("/message", handlers.NewMessage(db)).Methods(http.MethodPost)
	r.HandleFunc("/message", handlers.RoomMessages(db)).Queries("roomId", "{roomId}").Methods(http.MethodGet)
	r.HandleFunc("/message/after/{id}", handlers.RoomMessagesAfterMessage(db)).Queries("roomId", "{roomId}").Methods(http.MethodGet)
}
