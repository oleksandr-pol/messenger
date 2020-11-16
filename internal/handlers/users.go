package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/messenger/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
)

func CreateUser(db models.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&u); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		id, err := db.CreateUser(u)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, id)
	}
}

func UserRooms(db models.RoomsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		rooms, err := db.UserRooms(id)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "rooms not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
		}

		utils.RespondWithJSON(w, http.StatusOK, rooms)
	}
}
