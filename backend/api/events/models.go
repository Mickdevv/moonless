package events

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id               uuid.UUID `json:"id"`
	Type             string    `json:"type"`
	PosterPath       string    `json:"poster_path"`
	IsFeatured       bool      `json:"is_featured"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	Description      string    `json:"description"`
	Title            string    `json:"title"`
	LocationName     string    `json:"location_name"`
	LocationCity     string    `json:"location_city"`
	LocationMapsLink string    `json:"location_maps_link"`
	Active           bool      `json:"active"`
}
