
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


-- name: GetUserByEmail :one
select id, created_at, updated_at, email, email_verified_at, deactivated_at from users where email = $1;

-- name: GetUserById :one
select id, created_at, updated_at, email, email_verified_at, deactivated_at from users where id = $1;

-- name: GetUserByEmailForAuth :one
select id, created_at, updated_at, email, email_verified_at, deactivated_at, password from users where email = $1;
