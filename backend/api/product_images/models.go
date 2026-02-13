package productimages

import (
	"time"

	"github.com/google/uuid"
)

type ProductImagePayload struct {
	Id        uuid.UUID `json:"id"`
	ProductId uuid.UUID `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	IsPrimary bool      `json:"is_primary"`
}
