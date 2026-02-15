package productimages

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *utils.ServerCfg) {
	mux.HandleFunc("POST /product-images", CreateProductImage(serverCfg))
	mux.HandleFunc("DELETE /product-images/{id}", DeleteProductImage(serverCfg))
}
