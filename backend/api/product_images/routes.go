package productimages

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/auth"
	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("GET /api/product-images/{id}", GetProductImages(serverCfg))
	mux.HandleFunc("POST /api/product-images", auth.AuthMiddleware(serverCfg, CreateProductImage(serverCfg)))
	mux.HandleFunc("DELETE /api/product-images/{id}", auth.AuthMiddleware(serverCfg, DeleteProductImage(serverCfg)))
	mux.HandleFunc("PUT /api/product-images/{id}", auth.AuthMiddleware(serverCfg, MakeProductImagePrimary(serverCfg)))
}
