-- name: GetProductImagesByProductId :many
select * from product_images where product_id = $1;

-- name: GetProductImageByImageId :one
select * from product_images where id = $1;

-- name: CreateProductImage :one
insert into product_images(id, product_id, created_at, updated_at, path, is_primary)
values (
  gen_random_UUID(),
  $1,
  NOW(),
  NOW(),
  $2,
  $3
  )
  returning id, product_id, created_at, updated_at, path, is_primary;

-- name: DeleteProductImage :exec
delete from product_images where id = $1;

-- name: UpdateProductImage :exec
update product_images set updated_at = NOW(), is_primary = $2 where id = $1;

