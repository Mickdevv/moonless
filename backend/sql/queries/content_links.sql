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
  NOW(), -- created_at
  NOW() -- updated_at
  ) 
  returning id, platform, title, description, url, thumbnail_url, published_at, created_at, updated_at;

-- name: GetContentLinkById :one
SELECT id, platform, title, description, url, thumbnail_url, published_at, created_at, updated_at from content_links where id = $1;

-- name: GetContentLinks :many
SELECT id, platform, title, description, url, thumbnail_url, published_at, created_at, updated_at from content_links;
