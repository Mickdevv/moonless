package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
)

func LoginHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
	}
}
func RegisterHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Username  string `json:"username"`
			Password1 string `json:"password1"`
			Password2 string `json:"password2"`
		}

		data := parameters{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		errs := validatePassword(data.Password1)
		if errs != nil {
			utils.RespondWithJson(w, http.StatusBadRequest, struct {
				ValidationErrors []string `json:"validation_errors"`
			}{ValidationErrors: errs})
		}
		if data.Password1 != data.Password2 {
			utils.RespondWithError(w, http.StatusBadRequest, "Passwords do not match", err)
			return
		}

		serverCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Email:    data.Username,
			Password: "", // TODO : Create password hashing and checking logic
		})
	}
}

func validatePassword(password string) []string {
	var validationErrors []string
	if len(password) < 8 {
		validationErrors = append(validationErrors, "Password must be at least 8 characters long")
	}

	return validationErrors
}
