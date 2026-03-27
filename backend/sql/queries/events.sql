-- name: CreateEvent :one
INSERT INTO events(id, type, poster_path, is_featured, created_at, updated_at, start_date, end_date, description, title, location_name, location_city, location_maps_link)
VALUES (
  gen_random_UUID(),
  $1, -- type
  $2, -- poster_path
  $3, -- is_featured
  NOW(), -- created_at
  NOW(), -- updated_at
  $4, -- start_date
  $5, -- end_date
  $6, -- description
  $7, -- title
  $8, -- location_name
  $9, -- location_city
  $10 -- location_maps_link
  )
returning  id, type, poster_path, is_featured, created_at, updated_at, start_date, end_date, description, title, location_name, location_city, location_maps_link, active;

-- name: GetAllEvents :many
SELECT id, type, poster_path, is_featured, created_at, updated_at, start_date, end_date, description, title, location_name, location_city, location_maps_link, active from events;

-- name: GetActiveEvents :many
SELECT id, type, poster_path, is_featured, created_at, updated_at, start_date, end_date, description, title, location_name, location_city, location_maps_link, active from events where active = true;

-- name: GetEventById :one
SELECT id, type, poster_path, is_featured, created_at, updated_at, start_date, end_date, description, title, location_name, location_city, location_maps_link, active from events where id = $1;

