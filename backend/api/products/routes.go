package products

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("GET /products", GetProducts(serverCfg))
	mux.HandleFunc("POST /products", CreateProduct(serverCfg))

}
