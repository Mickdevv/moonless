package products

import (
	"time"

	"github.com/Mickdevv/moonless/backend/api/product_images"
	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID                           `json:"id"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   time.Time                           `json:"updated_at"`
	Active      bool                                `json:"active"`
	Description string                              `json:"description"`
	Name        string                              `json:"name"`
	Category    string                              `json:"category"`
	Price       int                                 `json:"price"`
	Stock       int                                 `json:"stock"`
	Images      []productimages.ProductImagePayload `json:"images"`
}

type UpdateProductPayload struct {
	Active      bool                                `json:"active"`
	Description string                              `json:"description"`
	Name        string                              `json:"name"`
	Category    string                              `json:"category"`
	Price       int                                 `json:"price"`
	Stock       int                                 `json:"stock"`
	Images      []productimages.ProductImagePayload `json:"images"`
}
type CreateProductPayload struct {
	Active      bool   `json:"active"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
