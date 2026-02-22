package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	productimages "github.com/Mickdevv/moonless/backend/api/product_images"
	"github.com/Mickdevv/moonless/backend/api/products"
	"github.com/Mickdevv/moonless/backend/api/utils"
	"github.com/Mickdevv/moonless/backend/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if os.Getenv("ENVIRONMENT") == "DEV" {
		dbURL += "?sslmode=disable"
	}
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	serverConfig := utils.ServerCfg{
		JWT_SECRET:       os.Getenv("JWT_SECRET"),
		DB:               database.New(dbConn),
		STATIC_FILES_DIR: os.Getenv("STATIC_FILES_DIR"),
	}

	mux := http.NewServeMux()

	fsHandler := http.StripPrefix("/api/static", http.FileServer(http.Dir(serverConfig.STATIC_FILES_DIR)))
	mux.Handle("/api/static/", fsHandler)

	products.RegisterRoutes(mux, &serverConfig)
	productimages.RegisterRoutes(mux, &serverConfig)

	serverPort := os.Getenv("SERVER_PORT")
	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: mux,
	}
	fmt.Println("Server listening on port", serverPort)
	log.Fatal(server.ListenAndServe())

}
