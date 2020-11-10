package handlers

import (
	"encoding/json"
	"net/http"

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
