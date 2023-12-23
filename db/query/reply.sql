-- name: CreateReply :one
INSERT INTO replies (
    tweet_id, user_id, content, updated_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetRepliesByTweetID :many
SELECT * FROM replies WHERE tweet_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: UpdateReply :one
UPDATE replies SET
    content = $2,
    updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteReply :exec
DELETE FROM replies WHERE id = $1;