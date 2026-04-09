package contentlinks

import (
	"time"

	"github.com/google/uuid"
)

type CreateContentLinkPayload struct {
	Platform    string    `json:"platform"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
}

type ContentLink struct {
	Id           uuid.UUID `json:"id"`
	Platform     string    `json:"platform"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	PublishedAt  time.Time `json:"published_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
