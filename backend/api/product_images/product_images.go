package productimages

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
	"github.com/google/uuid"
)

func CreateProductImage(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := r.FormValue("data")
		params := CreateProductImagePayload{}
		err := json.Unmarshal([]byte(data), &params)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "File error", err)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "File error", err)
			return
		}

		defer file.Close()

		filename := strconv.Itoa(int(time.Now().UnixMilli())) + filepath.Base(handler.Filename)

		destinationPath := filepath.Join(serverCfg.STATIC_FILES_DIR, "images", "products", filename)

		dst, err := os.Create(destinationPath)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error saving file", err)
			return
		}

		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error saving file", err)
			return
		}

		uploadedimage, err := serverCfg.DB.CreateProductImage(r.Context(), database.CreateProductImageParams{
			ProductID: params.ProductId,
			IsPrimary: params.IsPrimary,
			Path:      destinationPath,
		})

		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error saving file", err)
			return
		}

		type response struct {
			Message string              `json:"message"`
			Data    ProductImagePayload `json:"data"`
		}

		var res response

		res.Message = "Image uploaded successfully"

		res.Data.Id = uploadedimage.ID
		res.Data.CreatedAt = uploadedimage.CreatedAt
		res.Data.UpdatedAt = uploadedimage.UpdatedAt
		res.Data.ProductId = uploadedimage.ProductID
		res.Data.Path = uploadedimage.Path
		res.Data.IsPrimary = uploadedimage.IsPrimary

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func DeleteProductImage(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageIdStr := r.PathValue("id")
		imageId, err := uuid.Parse(imageIdStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid image id", err)
			return
		}

		image, err := serverCfg.DB.GetProductImageByImageId(r.Context(), imageId)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Image not found", err)
			return
		}

		err = serverCfg.DB.DeleteProductImage(r.Context(), imageId)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "An error occurred while deleting the image", err)
			return
		}

		// err := os.Remove(image.Path)

		type response struct {
			ProductImage ProductImagePayload `json:"product_image"`
		}

		res := response{
			ProductImage: ProductImagePayload{
				Id:        image.ID,
				Path:      image.Path,
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
			return
		}

		_, err = serverCfg.DB.GetProductById(r.Context(), productId)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Invalid product id", err)
			return
		}

		productImages, err := serverCfg.DB.GetProductImagesByProductId(r.Context(), productId)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error finding images", err)
			return
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
				Path:      image.Path,
				IsPrimary: image.IsPrimary,
				ProductId: image.ProductID,
			})
		}
		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
