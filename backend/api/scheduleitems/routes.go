package scheduleitems

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/auth"
	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("GET /api/schedule-items", GetAllScheduleItemsHandler(serverCfg))
	mux.HandleFunc("GET /api/schedule-items/{id}", GetScheduleItemByIdHandler(serverCfg))
	mux.HandleFunc("DELETE /api/schedule-items/{id}", DeleteScheduleItemByIdHandler(serverCfg))
	mux.HandleFunc("POST /api/schedule-items", auth.AuthMiddleware(serverCfg, CreateScheduleItemHandler(serverCfg)))
}
