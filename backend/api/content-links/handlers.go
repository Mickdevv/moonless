package contentlinks

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
	"github.com/google/uuid"
)

func CreateContentLinkHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := CreateContentLinkPayload{}

		err := json.Unmarshal([]byte(r.FormValue("data")), &params)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		filePath, err := utils.CreateStaticFile(serverCfg, filepath.Join("images", "products"), r)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error creating file", err)
			return
		}

		contentLink, err := serverCfg.DB.CreateContentLink(r.Context(), database.CreateContentLinkParams{
			Title:        params.Title,
			Description:  sql.NullString{String: params.Description},
			Platform:     params.Platform,
			Url:          params.Url,
			ThumbnailUrl: sql.NullString{String: filePath},
			PublishedAt:  sql.NullTime{Time: params.PublishedAt},
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Could not create content link", err)
			return
		}

		res := ContentLink{
			Title:        contentLink.Title,
			Description:  contentLink.Description.String,
			Platform:     contentLink.Platform,
			Url:          contentLink.Url,
			ThumbnailUrl: contentLink.ThumbnailUrl.String,
			PublishedAt:  contentLink.PublishedAt.Time,
			CreatedAt:    contentLink.CreatedAt.Time,
			UpdatedAt:    contentLink.UpdatedAt.Time,
		}

		utils.RespondWithJson(w, http.StatusOK, res)

	}
}

func GetContentLinkByIdHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Invalid link id", err)
			return
		}

		contentLink, err := serverCfg.DB.GetContentLinkById(r.Context(), id)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Could not retrieve content links", err)
			return
		}

		res := ContentLink{
			Title:        contentLink.Title,
			Description:  contentLink.Description.String,
			Platform:     contentLink.Platform,
			Url:          contentLink.Url,
			ThumbnailUrl: contentLink.ThumbnailUrl.String,
			PublishedAt:  contentLink.PublishedAt.Time,
			CreatedAt:    contentLink.CreatedAt.Time,
			UpdatedAt:    contentLink.UpdatedAt.Time,
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
func GetContentLinksHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentLinks, err := serverCfg.DB.GetContentLinks(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Could not retrieve content links", err)
			return
		}

		var res []ContentLink

		for _, contentLink := range contentLinks {
			res = append(res, ContentLink{
				Title:        contentLink.Title,
				Description:  contentLink.Description.String,
				Platform:     contentLink.Platform,
				Url:          contentLink.Url,
				ThumbnailUrl: contentLink.ThumbnailUrl.String,
				PublishedAt:  contentLink.PublishedAt.Time,
				CreatedAt:    contentLink.CreatedAt.Time,
				UpdatedAt:    contentLink.UpdatedAt.Time,
			})
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
