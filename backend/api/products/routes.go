package products

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/auth"
	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("GET /api/products", GetProducts(serverCfg))
	mux.HandleFunc("GET /api/products/{id}", GetProductById(serverCfg))
	mux.HandleFunc("POST /api/products", auth.AuthMiddleware(serverCfg, CreateProduct(serverCfg)))
	mux.HandleFunc("DELETE /api/products/{id}", auth.AuthMiddleware(serverCfg, DeleteProduct(serverCfg)))
	mux.HandleFunc("PUT /api/products/{id}", auth.AuthMiddleware(serverCfg, UpdateProduct(serverCfg)))
}
