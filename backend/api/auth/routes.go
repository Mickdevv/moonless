package auth

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("POST /api/register", RegisterHandler(serverCfg))
	mux.HandleFunc("POST /api/login", LoginHandler(serverCfg))
}
