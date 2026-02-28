-- name: MakeRefreshToken :one
INSERT INTO refresh_tokens(id, user_id, created_at, expires_at) 
  values (
gen_random_UUID(),
  $1,
  NOW(),
  $2
  )
  returning id, expires_at;

-- name: GetRefreshToken :one
SELECT id, expires_at, revoked_at, user_id, created_at from refresh_tokens where id = $1; 

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens SET revoked_at = NOW() WHERE id = $1;

