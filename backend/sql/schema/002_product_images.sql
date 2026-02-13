-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_images(
id UUID primary key unique,
  product_id uuid not null references products(id) on delete cascade,
  created_at timestamp not null default NOW(),
  updated_at timestamp not null default NOW(),
  url text unique not null,
  is_primary bool not null default false
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE product_images;
