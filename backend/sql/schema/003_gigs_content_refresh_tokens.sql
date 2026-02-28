-- +goose Up
-- +goose StatementBegin
CREATE TABLE events(
id UUID primary key unique,
  type text not null,
  poster_path text,
  is_featured boolean, 
  created_at timestamp not null default NOW(),
  updated_at timestamp not null default NOW(),
  start_date timestamp not null,
  end_date timestamp,
  description text,
  title text not null,
  location_name text,
  location_city text,
  location_maps_link text
);

CREATE TABLE content_links (
  id UUID PRIMARY KEY,
  platform text NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  url TEXT NOT NULL,
  thumbnail_url TEXT,
  published_at TIMESTAMP,
  is_featured BOOLEAN DEFAULT false,
  sort_order INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE refresh_tokens(
  id UUID primary key unique,
  user_id UUID not null ,
  foreign key (user_id) references users(id) on delete cascade,
  created_at timestamp not null default NOW(),
  expires_at timestamp not null,
  revoked_at timestamp
);

ALTER TABLE users ADD COLUMN role text not null default 'basic'; 
-- +goose StatementEnd

-- +goose Down 
ALTER TABLE users DROP COLUMN role;
DROP TABLE events;
DROP TABLE content_links;
DROP TABLE refresh_tokens;
