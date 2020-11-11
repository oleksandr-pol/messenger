package env

import (
	"database/sql"

	"github.com/oleksandr-pol/messenger/internal/routes"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/messenger/internal/models"
)

func SetUpServer(db *sql.DB) (*mux.Router, error) {
	router := mux.NewRouter()
	userDB := &models.UserDB{db}
	roomDB := &models.RoomDB{db}

	routes.User(router, userDB)
	routes.Room(router, roomDB)

	return router, nil
}
