package email

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func sendEmail(serverCfg *utils.ServerCfg, email SendEmail) error {
	auth := smtp.PlainAuth("", serverCfg.SMTP_FROM_ADDRESS, serverCfg.SMTP_PASSWORD, serverCfg.SMTP_HOST)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n\r\n%s",
		serverCfg.SMTP_FROM_ADDRESS,
		strings.Join(email.ToList, ", "),
		email.Subject,
		email.Msg,
	))

	return smtp.SendMail(
		serverCfg.SMTP_HOST+":"+serverCfg.SMTP_PORT,
		auth,
		serverCfg.SMTP_FROM_ADDRESS,
		email.ToList,
		msg,
	)
}
