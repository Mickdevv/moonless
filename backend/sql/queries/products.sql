-- name: CreateProduct :one
insert into products(id, created_at, updated_at, active, description, name, category, stock, price)
values (
  gen_random_UUID(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
  )
  returning id, created_at, updated_at, active, description, name, category, stock, price;

-- name: UpdateProduct :one
update products set updated_at = NOW(), name = $2, description = $3, category = $4, stock = $5, price = $6, active = $7 where id = $1 returning *; 

-- name: GetActiveProducts :many
select * from products where active = true;

-- name: GetAllProducts :many
select * from products;

-- name: GetProductById :one
select * from products where id = $1;

-- name: DeleteProduct :exec
delete from products where id = $1;
