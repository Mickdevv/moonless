package events

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
	"github.com/google/uuid"
)

func DeleteEventByIdHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid event id", err)
			return
		}

		err = serverCfg.DB.DeleteEventById(r.Context(), id)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Event not found", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
func CreateEventHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("CreateEventHandler")

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error parsing request", err)
			return
		}

		data := CreateEventPayload{}
		err = json.Unmarshal([]byte(r.FormValue("data")), &data)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		filePath, err := utils.CreateStaticFile(serverCfg, filepath.Join("images", "events"), r)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "File error", err)
			return
		}

		event, err := serverCfg.DB.CreateEvent(r.Context(), database.CreateEventParams{
			Type:             data.Type,
			PosterPath:       sql.NullString{String: filePath, Valid: true},
			IsFeatured:       data.IsFeatured,
			StartDate:        data.StartDate,
			EndDate:          sql.NullTime{Time: data.EndDate, Valid: true},
			Description:      data.Description,
			Title:            data.Title,
			LocationName:     sql.NullString{String: data.LocationName, Valid: true},
			LocationCity:     sql.NullString{String: data.LocationCity, Valid: true},
			LocationMapsLink: sql.NullString{String: data.LocationMapsLink, Valid: true},
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		res := Event{
			Id:               event.ID,
			Type:             event.Type,
			PosterPath:       event.PosterPath.String,
			IsFeatured:       event.IsFeatured,
			StartDate:        event.StartDate,
			EndDate:          event.EndDate.Time,
			Description:      event.Description,
			Title:            event.Title,
			LocationName:     event.LocationName.String,
			LocationCity:     event.LocationCity.String,
			LocationMapsLink: event.LocationMapsLink.String,
			Active:           event.Active.Bool,
			CreatedAt:        event.CreatedAt,
			UpdatedAt:        event.UpdatedAt,
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetActiveEventsHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		events, err := serverCfg.DB.GetActiveEvents(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Error fetching events", err)
			return
		}

		res := []Event{}

		for _, event := range events {
			res = append(res, Event{
				Id:               event.ID,
				Type:             event.Type,
				PosterPath:       event.PosterPath.String,
				IsFeatured:       event.IsFeatured,
				StartDate:        event.StartDate,
				EndDate:          event.EndDate.Time,
				Description:      event.Description,
				Title:            event.Title,
				LocationName:     event.LocationName.String,
				LocationCity:     event.LocationCity.String,
				LocationMapsLink: event.LocationMapsLink.String,
				Active:           event.Active.Bool,
				CreatedAt:        event.CreatedAt,
				UpdatedAt:        event.UpdatedAt,
			})
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetAllEventsHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		events, err := serverCfg.DB.GetAllEvents(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Error fetching events", err)
			return
		}

		res := []Event{}

		for _, event := range events {
			res = append(res, Event{
				Id:               event.ID,
				Type:             event.Type,
				PosterPath:       event.PosterPath.String,
				IsFeatured:       event.IsFeatured,
				StartDate:        event.StartDate,
				EndDate:          event.EndDate.Time,
				Description:      event.Description,
				Title:            event.Title,
				LocationName:     event.LocationName.String,
				LocationCity:     event.LocationCity.String,
				LocationMapsLink: event.LocationMapsLink.String,
				Active:           event.Active.Bool,
				CreatedAt:        event.CreatedAt,
				UpdatedAt:        event.UpdatedAt,
			})
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetEventByIdHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid event id", err)
			return
		}

		event, err := serverCfg.DB.GetEventById(r.Context(), id)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Event not found", err)
			return
		}

		res := Event{
			Id:               event.ID,
			Type:             event.Type,
			PosterPath:       event.PosterPath.String,
			IsFeatured:       event.IsFeatured,
			StartDate:        event.StartDate,
			EndDate:          event.EndDate.Time,
			Description:      event.Description,
			Title:            event.Title,
			LocationName:     event.LocationName.String,
			LocationCity:     event.LocationCity.String,
			LocationMapsLink: event.LocationMapsLink.String,
			Active:           event.Active.Bool,
			CreatedAt:        event.CreatedAt,
			UpdatedAt:        event.UpdatedAt,
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
