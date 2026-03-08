-- +goose Up
create table posts
(
  id           uuid primary key,
  created_at   timestamp    not null,
  updated_at   timestamp    not null,
  title        varchar(255) not null,
  url          varchar(255) unique not null,
  description  text,
  published_at timestamp,
  feed_id      uuid         not null references feeds (id) on delete cascade
);

-- +goose Down
drop table posts;
