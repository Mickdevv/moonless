package productimages

import (
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/google/uuid"
)

func DeleteProductImage(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageIdStr := r.PathValue("id")
		imageId, err := uuid.Parse(imageIdStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid image id", err)
		}

		image, err := serverCfg.DB.GetProductImageByImageId(r.Context(), imageId)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Image not found", err)
			return
		}

		type response struct {
			ProductImage ProductImagePayload `json:"product_image"`
		}

		res := response{
			ProductImage: ProductImagePayload{
				Id:        image.ID,
				Url:       image.Url,
				CreatedAt: image.CreatedAt,
				UpdatedAt: image.UpdatedAt,
				IsPrimary: image.IsPrimary,
				ProductId: image.ProductID,
			},
		}
		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetProductImages(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIdStr := r.PathValue("id")
		productId, err := uuid.Parse(productIdStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product id", err)
		}

		_, err = serverCfg.DB.GetProductById(r.Context(), productId)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Invalid product id", err)
		}

		productImages, err := serverCfg.DB.GetProductImagesByProductId(r.Context(), productId)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error finding images", err)
		}

		type response struct {
			Images []ProductImagePayload `json:"product_images"`
		}

		res := response{}
		for _, image := range productImages {
			res.Images = append(res.Images, ProductImagePayload{
				Id:        image.ID,
				UpdatedAt: image.UpdatedAt,
				CreatedAt: image.CreatedAt,
				Url:       image.Url,
				IsPrimary: image.IsPrimary,
				ProductId: image.ProductID,
			})
		}
		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
