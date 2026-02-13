
-- name: CreateUser :one
insert into users(id, created_at, updated_at, email, email_verified_at, password, deactivated_at)
values(
  gen_random_UUID(),
  NOW(),
  NOW(),
  $1,
  NULL,
  $2,
  NULL
  )
  returning id, created_at, updated_at, email, email_verified_at, deactivated_at;
