package events

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/auth"
	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("GET /api/events", GetAllEventsHandler(serverCfg))
	mux.HandleFunc("GET /api/events/{id}", GetEventByIdHandler(serverCfg))
	mux.HandleFunc("DELETE /api/events/{id}", DeleteEventByIdHandler(serverCfg))
	mux.HandleFunc("POST /api/events", auth.AuthMiddleware(serverCfg, CreateEventHandler(serverCfg)))
}
