-- name: CreateAccount :one
INSERT INTO accounts (
    username,
    email,
    profile_picture_url,
    bio,
    last_login,
    is_admin,
    is_active,
    last_deactivated_at,
    n_followers,
    n_following,
    n_tweets
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetAccountByUsername :one
SELECT * FROM accounts WHERE username = $1;

-- name: GetAccountByEmail :one
SELECT * FROM accounts WHERE email = $1;

-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = $1;

-- name: GetAccountByUsernameOrEmail :one
SELECT * FROM accounts WHERE username = $1 OR email = $2;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT $1 OFFSET $2;

-- name: AddNFollowersByOne :one
UPDATE accounts SET n_followers = n_followers + 1 WHERE id = $1 RETURNING *;

-- name: AddNFollowingByOne :one
UPDATE accounts SET n_following = n_following + 1 WHERE id = $1 RETURNING *;

-- name: AddNTweetsByOne :one
UPDATE accounts SET n_tweets = n_tweets + 1 WHERE id = $1 RETURNING *;

-- name: SubtractNFollowersByOne :one
UPDATE accounts SET n_followers = n_followers - 1 WHERE id = $1 RETURNING *;

-- name: SubtractNFollowingByOne :one
UPDATE accounts SET n_following = n_following - 1 WHERE id = $1 RETURNING *;

-- name: SubtractNTweetsByOne :one
UPDATE accounts SET n_tweets = n_tweets - 1 WHERE id = $1 RETURNING *;

-- name: SubtractAllNTweets :one
UPDATE accounts SET n_tweets = 0 WHERE id = $1 RETURNING *;

-- name: UpdateAccountByID :one
UPDATE accounts SET
    username = $2,
    email = $3,
    profile_picture_url = $4,
    bio = $5,
    last_login = $6,
    is_admin = $7,
    is_active = $8,
    last_deactivated_at = $9,
    n_followers = $10,
    n_following = $11,
    n_tweets = $12
WHERE id = $1
RETURNING *;

-- name: UpdateAccountByUsername :one
UPDATE accounts SET
    email = $2,
    profile_picture_url = $3,
    bio = $4,
    last_login = $5,
    is_admin = $6,
    is_active = $7,
    last_deactivated_at = $8,
    n_followers = $9,
    n_following = $10,
    n_tweets = $11
WHERE username = $1
RETURNING *;

-- name: UpdateAccountByEmail :one
UPDATE accounts SET
    username = $2,
    profile_picture_url = $3,
    bio = $4,
    last_login = $5,
    is_admin = $6,
    is_active = $7,
    last_deactivated_at = $8,
    n_followers = $9,
    n_following = $10,
    n_tweets = $11
WHERE email = $1 
RETURNING *;

-- name: DeleteAccountByID :exec
DELETE FROM accounts WHERE id = $1;

-- name: DeleteAccountByUsername :exec
DELETE FROM accounts WHERE username = $1;

-- name: DeleteAccountByEmail :exec
DELETE FROM accounts WHERE email = $1;