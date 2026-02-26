package productimages

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("POST /api/product-images", CreateProductImage(serverCfg))
	mux.HandleFunc("DELETE /api/product-images/{id}", DeleteProductImage(serverCfg))
	mux.HandleFunc("GET /api/product-images/{id}", GetProductImages(serverCfg))
	mux.HandleFunc("PUT /api/product-images/{id}", MakeProductImagePrimary(serverCfg))
}
