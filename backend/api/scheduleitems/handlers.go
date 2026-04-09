package scheduleitems

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
	"github.com/google/uuid"
)

func DeleteScheduleItemByIdHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid schedule item id", err)
			return
		}

		err = serverCfg.DB.DeleteEventById(r.Context(), id)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Schedule item not found", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
func CreateScheduleItemHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error parsing request", err)
			return
		}

		data := CreateScheduleItemPayload{}
		err = json.Unmarshal([]byte(r.FormValue("data")), &data)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
			return
		}

		filePath, err := utils.CreateStaticFile(serverCfg, filepath.Join("images", "schedule-items"), r)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "File error", err)
			return
		}

		scheduleItem, err := serverCfg.DB.CreateEvent(r.Context(), database.CreateEventParams{
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

		res := ScheduleItem{
			Id:               scheduleItem.ID,
			Type:             scheduleItem.Type,
			PosterPath:       scheduleItem.PosterPath.String,
			IsFeatured:       scheduleItem.IsFeatured,
			StartDate:        scheduleItem.StartDate,
			EndDate:          scheduleItem.EndDate.Time,
			Description:      scheduleItem.Description,
			Title:            scheduleItem.Title,
			LocationName:     scheduleItem.LocationName.String,
			LocationCity:     scheduleItem.LocationCity.String,
			LocationMapsLink: scheduleItem.LocationMapsLink.String,
			Active:           scheduleItem.Active.Bool,
			CreatedAt:        scheduleItem.CreatedAt,
			UpdatedAt:        scheduleItem.UpdatedAt,
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetActiveScheduleItemsHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		scheduleItems, err := serverCfg.DB.GetActiveEvents(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Error fetching schedule items", err)
			return
		}

		res := []ScheduleItem{}

		for _, scheduleItem := range scheduleItems {
			res = append(res, ScheduleItem{
				Id:               scheduleItem.ID,
				Type:             scheduleItem.Type,
				PosterPath:       scheduleItem.PosterPath.String,
				IsFeatured:       scheduleItem.IsFeatured,
				StartDate:        scheduleItem.StartDate,
				EndDate:          scheduleItem.EndDate.Time,
				Description:      scheduleItem.Description,
				Title:            scheduleItem.Title,
				LocationName:     scheduleItem.LocationName.String,
				LocationCity:     scheduleItem.LocationCity.String,
				LocationMapsLink: scheduleItem.LocationMapsLink.String,
				Active:           scheduleItem.Active.Bool,
				CreatedAt:        scheduleItem.CreatedAt,
				UpdatedAt:        scheduleItem.UpdatedAt,
			})
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetAllScheduleItemsHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		scheduleItems, err := serverCfg.DB.GetAllEvents(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Error fetching schedule items", err)
			return
		}

		res := []ScheduleItem{}

		for _, scheduleItems := range scheduleItems {
			res = append(res, ScheduleItem{
				Id:               scheduleItems.ID,
				Type:             scheduleItems.Type,
				PosterPath:       scheduleItems.PosterPath.String,
				IsFeatured:       scheduleItems.IsFeatured,
				StartDate:        scheduleItems.StartDate,
				EndDate:          scheduleItems.EndDate.Time,
				Description:      scheduleItems.Description,
				Title:            scheduleItems.Title,
				LocationName:     scheduleItems.LocationName.String,
				LocationCity:     scheduleItems.LocationCity.String,
				LocationMapsLink: scheduleItems.LocationMapsLink.String,
				Active:           scheduleItems.Active.Bool,
				CreatedAt:        scheduleItems.CreatedAt,
				UpdatedAt:        scheduleItems.UpdatedAt,
			})
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}

func GetScheduleItemByIdHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid schedule item id", err)
			return
		}

		scheduleItem, err := serverCfg.DB.GetEventById(r.Context(), id)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Schedule item not found", err)
			return
		}

		res := ScheduleItem{
			Id:               scheduleItem.ID,
			Type:             scheduleItem.Type,
			PosterPath:       scheduleItem.PosterPath.String,
			IsFeatured:       scheduleItem.IsFeatured,
			StartDate:        scheduleItem.StartDate,
			EndDate:          scheduleItem.EndDate.Time,
			Description:      scheduleItem.Description,
			Title:            scheduleItem.Title,
			LocationName:     scheduleItem.LocationName.String,
			LocationCity:     scheduleItem.LocationCity.String,
			LocationMapsLink: scheduleItem.LocationMapsLink.String,
			Active:           scheduleItem.Active.Bool,
			CreatedAt:        scheduleItem.CreatedAt,
			UpdatedAt:        scheduleItem.UpdatedAt,
		}

		utils.RespondWithJson(w, http.StatusOK, res)
	}
}
