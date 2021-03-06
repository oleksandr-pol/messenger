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
	participantDB := &models.ParticipantDB{db}
	messageDB := &models.MessageDB{db}

	routes.User(router, userDB, roomDB)
	routes.Room(router, roomDB)
	routes.Participant(router, participantDB)
	routes.Message(router, messageDB)

	return router, nil
}
