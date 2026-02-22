package productimages

import (
	"time"

	"github.com/google/uuid"
)

type CreateProductImagePayload struct {
	ProductId uuid.UUID `json:"product_id"`
	IsPrimary bool      `json:"is_primary"`
}
type ProductImagePayload struct {
	Id        uuid.UUID `json:"id"`
	ProductId uuid.UUID `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Path      string    `json:"path"`
	IsPrimary bool      `json:"is_primary"`
}
