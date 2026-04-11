package email

import (
	"net/smtp"

	"github.com/Mickdevv/moonless/backend/api/utils"
)

func sendEmail(serverCfg *utils.ServerCfg, email SendEmail) error {
	// server config will have smtp srv port and host

	body := []byte(email.Msg)

	smtpConnection := smtp.PlainAuth("", serverCfg.SMTP_FROM_ADDRESS, serverCfg.SMTP_PASSWORD, serverCfg.SMTP_HOST)

	return smtp.SendMail(serverCfg.SMTP_HOST+":"+serverCfg.SMTP_PORT, smtpConnection, serverCfg.SMTP_FROM_ADDRESS, email.ToList, body)

}
