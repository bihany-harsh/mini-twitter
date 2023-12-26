-- name: CreateFollow :one
INSERT INTO follows (
    follower_id, following_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetFollowers :many
SELECT * FROM follows WHERE following_id = $1 ORDER BY follower_id LIMIT $2 OFFSET $3;

-- name: GetFollowing :many
SELECT * FROM follows WHERE follower_id = $1 ORDER BY following_id LIMIT $2 OFFSET $3;

-- name: GetFollow :one
SELECT * FROM follows WHERE follower_id = $1 AND following_id = $2;

-- name: DeleteFollow :exec
DELETE FROM follows WHERE follower_id = $1 AND following_id = $2;