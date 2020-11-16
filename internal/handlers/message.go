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

func NewMessage(db models.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var message models.Message

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&message); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		id, err := db.CreateMessage(message)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, id)
	}
}

func RoomMessages(db models.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("roomId")
		roomId, err := strconv.Atoi(id)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
			return
		}

		messages, err := db.RoomMessages(roomId)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "messages not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
		}

		utils.RespondWithJSON(w, http.StatusOK, messages)
	}
}

func RoomMessagesAfterMessage(db models.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rId := r.FormValue("roomId")
		roomId, err := strconv.Atoi(rId)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
			return
		}

		vars := mux.Vars(r)
		msgId, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid message ID")
			return
		}

		messages, err := db.RoomMessagesAfter(roomId, msgId)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "messages not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
		}

		utils.RespondWithJSON(w, http.StatusOK, messages)
	}
}
