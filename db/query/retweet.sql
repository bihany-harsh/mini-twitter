-- name: CreateRetweet :one
INSERT INTO retweets (
    user_id, tweet_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetRetweetsByUserID :many
SELECT * FROM retweets WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: GetRetweetsByTweetID :many
SELECT * FROM retweets WHERE tweet_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: DeleteRetweet :exec
DELETE FROM retweets WHERE user_id = $1 AND tweet_id = $2;