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

	routes.User(router, userDB)

	return router, nil
}
