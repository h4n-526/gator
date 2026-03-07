-- +goose Up
create table users
(
  id         uuid primary key,
  created_at timestamp           not null,
  updated_at timestamp           not null,
  name       varchar(255) unique not null
);

-- +goose Down
drop table users;
