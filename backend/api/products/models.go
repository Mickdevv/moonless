package products

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Active      bool      `json:"active"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
}

type CreateProductPayload struct {
	Active      bool   `json:"active"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
