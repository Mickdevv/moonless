package contentlinks

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/auth"
	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("GET /api/content-links", GetContentLinksHandler(serverCfg))
	mux.HandleFunc("GET /api/content-links/{id}", GetContentLinkByIdHandler(serverCfg))
	mux.HandleFunc("POST /api/content-links", auth.AuthMiddleware(serverCfg, CreateContentLinkHandler(serverCfg)))
	mux.HandleFunc("DELETE /api/content-links/{id}", auth.AuthMiddleware(serverCfg, DeleteContentLinkByIdHandler(serverCfg)))
}
