package contentlinks

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/auth"
	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("GET /api/discography", GetContentLinksHandler(serverCfg))
	mux.HandleFunc("GET /api/discography/{id}", GetContentLinkByIdHandler(serverCfg))
	mux.HandleFunc("POST /api/discography", auth.AuthMiddleware(serverCfg, CreateContentLinkHandler(serverCfg)))
	mux.HandleFunc("DELETE /api/discography/{id}", auth.AuthMiddleware(serverCfg, DeleteContentLinkByIdHandler(serverCfg)))
}
