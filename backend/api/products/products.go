package products

import (
	"encoding/json"
	"net/http"

	productimages "github.com/Mickdevv/moonless/backend/api/product_images"
	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
	"github.com/google/uuid"
)

func UpdateProduct(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Invalid product id", err)
			return
		}

		data := UpdateProductPayload{}
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&data)

		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		product, err := serverCfg.DB.UpdateProduct(r.Context(), database.UpdateProductParams{
			ID:          id,
			Active:      data.Active,
			Name:        data.Name,
			Description: data.Description,
			Price:       int32(data.Price),
			Category:    data.Category,
			Stock:       int32(data.Stock),
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error updating product", err)
			return
		}

		res := Product{
			Id:          product.ID,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Active:      product.Active,
			Price:       int(product.Price),
			Name:        product.Name,
			Description: product.Description,
			Images:      data.Images,
			Stock:       int(product.Stock),
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

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
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating product", err)
			return
		}

		res := Product{
			Id:          product.ID,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Active:      product.Active,
			Price:       int(product.Price),
			Name:        product.Name,
			Description: product.Description,
			Stock:       int(product.Stock),
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetProductById(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := uuid.Parse(idString)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Invalid id", err)
		}
		product, err := serverCfg.DB.GetProductById(r.Context(), id)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Product not found", err)
			return
		}

		type response struct {
			Data Product `json:"data"`
		}

		res := response{}
		productImages, err := serverCfg.DB.GetProductImagesByProductId(r.Context(), product.ID)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Images not found for one or more products", err)
		}
		var images []productimages.ProductImagePayload
		for _, i := range productImages {
			images = append(images, productimages.ProductImagePayload{
				Id:        i.ID,
				ProductId: i.ProductID,
				CreatedAt: i.CreatedAt,
				UpdatedAt: i.UpdatedAt,
				Path:      i.Path,
				IsPrimary: i.IsPrimary,
			})
		}
		res.Data = Product{
			Id:          product.ID,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Active:      product.Active,
			Price:       int(product.Price),
			Name:        product.Name,
			Description: product.Description,
			Images:      images,
			Category:    product.Category,
			Stock:       int(product.Stock),
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
			productImages, err := serverCfg.DB.GetProductImagesByProductId(r.Context(), p.ID)
			if err != nil {
				utils.RespondWithError(w, http.StatusNotFound, "Images not found for one or more products", err)
			}
			var images []productimages.ProductImagePayload
			for _, i := range productImages {
				images = append(images, productimages.ProductImagePayload{
					Id:        i.ID,
					ProductId: i.ProductID,
					CreatedAt: i.CreatedAt,
					UpdatedAt: i.UpdatedAt,
					Path:      i.Path,
					IsPrimary: i.IsPrimary,
				})
			}
			res.Products = append(res.Products, Product{
				Id:          p.ID,
				CreatedAt:   p.CreatedAt,
				UpdatedAt:   p.UpdatedAt,
				Active:      p.Active,
				Price:       int(p.Price),
				Name:        p.Name,
				Description: p.Description,
				Images:      images,
				Category:    p.Category,
				Stock:       int(p.Stock),
			})
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
