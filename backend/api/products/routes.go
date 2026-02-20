package products

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("GET /api/products", GetProducts(serverCfg))
	mux.HandleFunc("POST /api/products", CreateProduct(serverCfg))

}
