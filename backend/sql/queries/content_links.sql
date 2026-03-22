-- name: CreateContentLink :one
INSERT INTO content_links(id, platform, title, description, url, thumbnail_url, published_at, created_at, updated_at) 
VALUES (
  gen_random_UUID(),
  $1, -- platform
  $2, -- title
  $3, -- description
  $4, -- url
  $5, -- thumbnail_url
  $6, -- published_at
  $7, -- created_at
  $8 -- updated_at
  ) 
  returning id, platform, title, description, url, thumbnail_url, published_at, created_at, updated_at;

-- name: GetContentLinks :many
SELECT id, platform, title, description, url, thumbnail_url, published_at, created_at, updated_at from content_links;
