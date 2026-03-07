-- name: CreateFeedFollow :one
with inserted_feed_follow as (
insert
into feed_follows (id, created_at, updated_at, user_id, feed_id)
values ($1, $2, $3, $4, $5)
  returning *
  )
select inserted_feed_follow.*,
       users.name as user_name,
       feeds.name as feed_name
from inserted_feed_follow
       join users on users.id = inserted_feed_follow.user_id
       join feeds on feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
select feed_follows.id,
       feed_follows.created_at,
       feed_follows.updated_at,
       feed_follows.user_id,
       feed_follows.feed_id,
       users.name  as user_name,
       feeds.name  as feed_name
from feed_follows
join users on users.id = feed_follows.user_id
join feeds on feeds.id = feed_follows.feed_id
where feed_follows.user_id = $1;
