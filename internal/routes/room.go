
package routes

import (
"net/http"

"github.com/gorilla/mux"
"github.com/oleksandr-pol/messenger/internal/handlers"
"github.com/oleksandr-pol/messenger/internal/models"
)

func Room(r *mux.Router, db models.RoomsRepository) {
	r.HandleFunc("/room", handlers.CreateRoom(db)).Methods(http.MethodPost)
}
