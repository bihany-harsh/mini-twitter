// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: tweet.sql

package db

import (
	"context"
	"database/sql"
)

const addNLikesByOne = `-- name: AddNLikesByOne :one
UPDATE tweets SET n_likes = n_likes + 1 WHERE id = $1 RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

func (q *Queries) AddNLikesByOne(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, addNLikesByOne, id)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const addNReplyByOne = `-- name: AddNReplyByOne :one
UPDATE tweets SET n_reply = n_reply + 1 WHERE id = $1 RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

func (q *Queries) AddNReplyByOne(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, addNReplyByOne, id)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const addNRetweetsByOne = `-- name: AddNRetweetsByOne :one
UPDATE tweets SET n_retweets = n_retweets + 1 WHERE id = $1 RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

func (q *Queries) AddNRetweetsByOne(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, addNRetweetsByOne, id)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const createTweet = `-- name: CreateTweet :one
INSERT INTO tweets (
    user_id, content, updated_at, retweet_id, n_likes, n_retweets, n_reply
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

type CreateTweetParams struct {
	UserID    int64         `json:"user_id"`
	Content   string        `json:"content"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	RetweetID sql.NullInt64 `json:"retweet_id"`
	NLikes    int32         `json:"n_likes"`
	NRetweets int32         `json:"n_retweets"`
	NReply    int32         `json:"n_reply"`
}

func (q *Queries) CreateTweet(ctx context.Context, arg CreateTweetParams) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, createTweet,
		arg.UserID,
		arg.Content,
		arg.UpdatedAt,
		arg.RetweetID,
		arg.NLikes,
		arg.NRetweets,
		arg.NReply,
	)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const deleteTweetByID = `-- name: DeleteTweetByID :exec
DELETE FROM tweets WHERE id = $1
`

func (q *Queries) DeleteTweetByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTweetByID, id)
	return err
}

const deleteTweetsByUserID = `-- name: DeleteTweetsByUserID :exec
DELETE FROM tweets WHERE user_id = $1
`

func (q *Queries) DeleteTweetsByUserID(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteTweetsByUserID, userID)
	return err
}

const getTweetByID = `-- name: GetTweetByID :one
SELECT id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply FROM tweets WHERE id = $1
`

func (q *Queries) GetTweetByID(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, getTweetByID, id)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const getTweetsByUserID = `-- name: GetTweetsByUserID :many
SELECT id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply FROM tweets WHERE user_id = $1 ORDER BY id LIMIT $2 OFFSET $3
`

type GetTweetsByUserIDParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetTweetsByUserID(ctx context.Context, arg GetTweetsByUserIDParams) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, getTweetsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tweet
	for rows.Next() {
		var i Tweet
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
			&i.RetweetID,
			&i.NLikes,
			&i.NRetweets,
			&i.NReply,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTweets = `-- name: ListTweets :many
SELECT id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply FROM tweets ORDER BY id LIMIT $1 OFFSET $2
`

type ListTweetsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTweets(ctx context.Context, arg ListTweetsParams) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, listTweets, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tweet
	for rows.Next() {
		var i Tweet
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
			&i.RetweetID,
			&i.NLikes,
			&i.NRetweets,
			&i.NReply,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setRetweetID = `-- name: SetRetweetID :one
UPDATE tweets SET retweet_id = $2 WHERE id = $1 RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

type SetRetweetIDParams struct {
	ID        int64         `json:"id"`
	RetweetID sql.NullInt64 `json:"retweet_id"`
}

func (q *Queries) SetRetweetID(ctx context.Context, arg SetRetweetIDParams) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, setRetweetID, arg.ID, arg.RetweetID)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const subtractNLikesByOne = `-- name: SubtractNLikesByOne :one
UPDATE tweets SET n_likes = n_likes - 1 WHERE id = $1 RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

func (q *Queries) SubtractNLikesByOne(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, subtractNLikesByOne, id)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const subtractNReplyByOne = `-- name: SubtractNReplyByOne :one
UPDATE tweets SET n_reply = n_reply - 1 WHERE id = $1 RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

func (q *Queries) SubtractNReplyByOne(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, subtractNReplyByOne, id)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const subtractNRetweetsByOne = `-- name: SubtractNRetweetsByOne :one
UPDATE tweets SET n_retweets = n_retweets - 1 WHERE id = $1 RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

func (q *Queries) SubtractNRetweetsByOne(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, subtractNRetweetsByOne, id)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}

const updateTweetByID = `-- name: UpdateTweetByID :one
UPDATE tweets SET
    content = $2,
    updated_at = $3,
    retweet_id = $4,
    n_likes = $5,
    n_retweets = $6,
    n_reply = $7
WHERE id = $1
RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

type UpdateTweetByIDParams struct {
	ID        int64         `json:"id"`
	Content   string        `json:"content"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	RetweetID sql.NullInt64 `json:"retweet_id"`
	NLikes    int32         `json:"n_likes"`
	NRetweets int32         `json:"n_retweets"`
	NReply    int32         `json:"n_reply"`
}

func (q *Queries) UpdateTweetByID(ctx context.Context, arg UpdateTweetByIDParams) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, updateTweetByID,
		arg.ID,
		arg.Content,
		arg.UpdatedAt,
		arg.RetweetID,
		arg.NLikes,
		arg.NRetweets,
		arg.NReply,
	)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.RetweetID,
		&i.NLikes,
		&i.NRetweets,
		&i.NReply,
	)
	return i, err
}
