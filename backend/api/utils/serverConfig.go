package utils

import "github.com/Mickdevv/moonless/backend/internal/database"

type ServerCfg struct {
	JWT_SECRET       string
	DB               *database.Queries
	STATIC_FILES_DIR string

	//Email Configuration
	SMTP_HOST         string
	SMTP_PORT         string
	SMTP_FROM_ADDRESS string
	SMTP_PASSWORD     string
	ADMIN_EMAIL       string
}
