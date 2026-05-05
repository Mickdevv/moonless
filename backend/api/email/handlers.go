package email

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func ContactFormEmailHandler(serverCfg *utils.ServerCfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := ContactFormEmailPayload{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "payload error", err)
			return
		}

		confirmation := SendEmail{
			ToList:  []string{params.Email},
			Subject: "We received your message – Moonless",
			Msg: fmt.Sprintf(
				"Hi %s,\r\n\r\nThanks for reaching out! We've received your message and will get back to you soon.\r\n\r\nYour message:\r\n%s\r\n\r\n– Moonless",
				params.Name,
				params.Message,
			),
		}
		if err := sendEmail(serverCfg, confirmation); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to send confirmation email", err)
			return
		}

		adminNotification := SendEmail{
			ToList:  []string{serverCfg.ADMIN_EMAIL},
			Subject: fmt.Sprintf("New contact form submission from %s", params.Name),
			Msg: fmt.Sprintf(
				"Name: %s\r\nEmail: %s\r\nPhone: %s\r\nReason: %s\r\n\r\nMessage:\r\n%s",
				params.Name,
				params.Email,
				params.PhoneNumber,
				params.Reason,
				params.Message,
			),
		}
		if err := sendEmail(serverCfg, adminNotification); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to send admin notification", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
