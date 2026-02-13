#!/bin/bash

docker-compose down -v
docker-compose up -d
sleep 4

godotenv -f ./.env goose up
sqlc generate
