package utils

import "github.com/Mickdevv/moonless/backend/internal/database"

type ServerCfg struct {
	JWT_SECRET       string
	DB               *database.Queries
	STATIC_FILES_DIR string
}
