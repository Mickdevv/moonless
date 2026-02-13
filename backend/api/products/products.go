package products

import (
	"encoding/json"
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
)

func CreateProduct(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := CreateProductPayload{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)

		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		product, err := serverCfg.DB.CreateProduct(r.Context(), database.CreateProductParams{
			Active:      data.Active,
			Name:        data.Name,
			Description: data.Description,
			Price:       int32(data.Price),
			Category:    data.Category,
			Stock:       int32(data.Stock),
		})

		res := Product{
			Id:          product.ID,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Active:      product.Active,
			Price:       int(product.Price),
			Name:        product.Name,
			Description: product.Description,
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetProducts(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := serverCfg.DB.GetAllProducts(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Products not found", err)
			return
		}

		type response struct {
			Products []Product `json:"products"`
		}

		res := response{}
		for _, p := range products {
			res.Products = append(res.Products, Product{
				Id:          p.ID,
				CreatedAt:   p.CreatedAt,
				UpdatedAt:   p.UpdatedAt,
				Active:      p.Active,
				Price:       int(p.Price),
				Name:        p.Name,
				Description: p.Description,
			})
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
