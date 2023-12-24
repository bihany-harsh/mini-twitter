-- name: CreateTweet :one
INSERT INTO tweets (
    user_id, content, updated_at, retweet_id, n_likes, n_retweets, n_reply
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetTweetByID :one
SELECT * FROM tweets WHERE id = $1;

-- name: GetTweetsByUserID :many
SELECT * FROM tweets WHERE user_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: ListTweets :many
SELECT * FROM tweets ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateTweetByID :one
UPDATE tweets SET
    content = $2,
    updated_at = $3,
    retweet_id = $4,
    n_likes = $5,
    n_retweets = $6,
    n_reply = $7
WHERE id = $1
RETURNING *;

-- name: DeleteTweetByID :exec
DELETE FROM tweets WHERE id = $1;

-- name: DeleteTweetsByUserID :exec
DELETE FROM tweets WHERE user_id = $1;