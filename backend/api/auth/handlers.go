package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RefreshTokenHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid id", err)
			return
		}

		refreshToken, err := serverCfg.DB.GetRefreshToken(r.Context(), id)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Refresh token not found or invalid", err)
			return
		}

		if refreshToken.ExpiresAt.Unix() < time.Now().Unix() || refreshToken.RevokedAt.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "Refresh token not found or invalid", nil)
			return
		}

		user, err := serverCfg.DB.GetUserById(r.Context(), refreshToken.UserID)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Refresh token not found or invalid", err)
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

		newRefreshTokenExpires := time.Now().Add(time.Minute * 60)
		newRefreshToken, err := serverCfg.DB.MakeRefreshToken(r.Context(), database.MakeRefreshTokenParams{
			UserID:    user.ID,
			ExpiresAt: newRefreshTokenExpires,
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Something went wrong. Please try again later.", err)
			return
		}

		err = serverCfg.DB.RevokeRefreshToken(r.Context(), refreshToken.ID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Something went wrong. Please try again later.", err)
			return
		}

		res := loginResponse{RefreshToken: newRefreshToken.ID.String(), RefreshTokenExpires: newRefreshTokenExpires, AccessToken: token}
		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
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
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(time.Minute * 70),
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

		utils.RespondWithJson(w, http.StatusOK, loginResponse{AccessToken: token, RefreshToken: refresh.ID.String(), RefreshTokenExpires: refresh.ExpiresAt})

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
