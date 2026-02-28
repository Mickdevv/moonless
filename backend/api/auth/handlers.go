package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data loginPayload

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		err := decoder.Decode(&data)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Incorrect payload", err)
			return
		}

		user, err := serverCfg.DB.GetUserByEmailForAuth(r.Context(), data.Email)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Invalid login credentials", err)
			return
		}

		ok, err := checkPassword(data.Password, user.Password)
		if !ok || err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Invalid login credentials", err)
			return
		}

		refresh, err := serverCfg.DB.MakeRefreshToken(r.Context(), database.MakeRefreshTokenParams{
			UserID: user.ID,
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Something went wrong. Please try again later.", err)
			return
		}

		token, err := makeJWT(serverCfg, CustomJwtClaims{
			Role:  user.Role,
			Email: user.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   user.ID.String(),
				Issuer:    "moonless",
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			},
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating token", err)
			return
		}

		utils.RespondWithJson(w, http.StatusOK, loginResponse{AccessToken: token, RefreshToken: refresh.ID.String()})

	}
}
func RegisterHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := registerPayload{}
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
			return
		}
		if data.Password1 != data.Password2 {
			utils.RespondWithError(w, http.StatusBadRequest, "Passwords do not match", err)
			return
		}
		if errors := validatePassword(data.Password1); len(errors) > 0 {
			utils.RespondWithJson(w, http.StatusBadRequest, errors)
			return
		}

		hashed_password, err := makePassword(data.Password1)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
			return
		}
		_, err = serverCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Email:    data.Email,
			Password: hashed_password,
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Username unavailable", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func validatePassword(password string) []string {
	var validationErrors []string
	if len(password) < 8 {
		validationErrors = append(validationErrors, "Password must be at least 8 characters long")
	}

	return validationErrors
}
