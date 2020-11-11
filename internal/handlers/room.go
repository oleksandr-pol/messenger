package handlers

import (
	"encoding/json"
	"github.com/oleksandr-pol/messenger/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
	"net/http"
)

func CreateRoom(db models.RoomsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var room models.Room

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&room); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		id, err := db.CreateRoom(room)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, id)
	}
}
