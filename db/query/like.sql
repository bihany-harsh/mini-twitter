-- name: CreateLike :one
INSERT INTO likes (
    user_id, tweet_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetLike :one
SELECT * FROM likes WHERE user_id = $1 AND tweet_id = $2;

-- name: GetLikesByUserID :many
SELECT * FROM likes WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: GetLikesByTweetID :many
SELECT * FROM likes WHERE tweet_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: DeleteLike :exec
DELETE FROM likes WHERE user_id = $1 AND tweet_id = $2;