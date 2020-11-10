package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/messenger/internal/handlers"
	"github.com/oleksandr-pol/messenger/internal/models"
)

func User(r *mux.Router, db models.UsersRepository) {
	r.HandleFunc("/user", handlers.CreateUser(db)).Methods(http.MethodPost)
}
