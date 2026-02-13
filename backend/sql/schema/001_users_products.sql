-- +goose Up
-- +goose StatementBegin

CREATE TABLE users(
id UUID primary key unique,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  email text NOT NULL UNIQUE,
  email_verified_at timestamp DEFAULT NULL,
  password text NOT NULL DEFAULT 'unset',
  deactivated_at timestamp DEFAULT NULL
);

CREATE TABLE products(
  id UUID primary key unique,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  active bool NOT NULL default true,
  description text not null default '',
  name text NOT NULL,
  category text NOT NULL,
  price int not null default 1,
  stock int NOT NULL default 0
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE products;
-- +goose StatementEnd
