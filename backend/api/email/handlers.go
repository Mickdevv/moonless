package email

import (
	"encoding/json"
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func ContactFormEmailHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := ContactFormEmailPayload{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&params)

		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "payload error", err)
			return
		}

		email := SendEmail{
			ToList:  []string{params.Email},
			Subject: "hey dude",
			Msg:     params.Message,
		}

		err = sendEmail(serverCfg, email)

		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "email error", err)
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}
