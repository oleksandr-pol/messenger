package handlers

import (
	"encoding/json"
	"net/http"

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
