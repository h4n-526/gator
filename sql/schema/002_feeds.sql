-- +goose Up
create table feeds
(
  id         uuid primary key,
  created_at timestamp           not null,
  updated_at timestamp           not null,
  name       varchar(255)        not null,
  url        varchar(255) unique not null,
  user_id    uuid                not null references users (id) on delete cascade
);

-- +goose Down
drop table feeds;
