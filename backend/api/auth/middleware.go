package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func AuthMiddleware(serverCfg *utils.ServerCfg, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		tokenStr := strings.TrimPrefix(tokenHeader, "Bearer ")

		claims, err := validateJWT(serverCfg, tokenStr)
		if err != nil {

			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorised", err)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
