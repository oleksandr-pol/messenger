package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/messenger/internal/handlers"
	"github.com/oleksandr-pol/messenger/internal/models"
)

func User(r *mux.Router, userDb models.UsersRepository, roomDb models.RoomsRepository) {
	r.HandleFunc("/user", handlers.CreateUser(userDb)).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}/rooms", handlers.UserRooms(roomDb)).Methods(http.MethodGet)
}
