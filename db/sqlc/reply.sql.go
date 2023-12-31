// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: reply.sql

package db

import (
	"context"
	"database/sql"
)

const createReply = `-- name: CreateReply :one
INSERT INTO replies (
    tweet_id, user_id, content, updated_at
) VALUES (
    $1, $2, $3, $4
) RETURNING id, tweet_id, user_id, content, created_at, updated_at
`

type CreateReplyParams struct {
	TweetID   int64        `json:"tweet_id"`
	UserID    int64        `json:"user_id"`
	Content   string       `json:"content"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateReply(ctx context.Context, arg CreateReplyParams) (Reply, error) {
	row := q.db.QueryRowContext(ctx, createReply,
		arg.TweetID,
		arg.UserID,
		arg.Content,
		arg.UpdatedAt,
	)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.TweetID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteReply = `-- name: DeleteReply :exec
DELETE FROM replies WHERE id = $1
`

func (q *Queries) DeleteReply(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReply, id)
	return err
}

const getRepliesByTweetID = `-- name: GetRepliesByTweetID :many
SELECT id, tweet_id, user_id, content, created_at, updated_at FROM replies WHERE tweet_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3
`

type GetRepliesByTweetIDParams struct {
	TweetID int64 `json:"tweet_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) GetRepliesByTweetID(ctx context.Context, arg GetRepliesByTweetIDParams) ([]Reply, error) {
	rows, err := q.db.QueryContext(ctx, getRepliesByTweetID, arg.TweetID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reply
	for rows.Next() {
		var i Reply
		if err := rows.Scan(
			&i.ID,
			&i.TweetID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getRepliesByUserID = `-- name: GetRepliesByUserID :many
SELECT id, tweet_id, user_id, content, created_at, updated_at FROM replies WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3
`

type GetRepliesByUserIDParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetRepliesByUserID(ctx context.Context, arg GetRepliesByUserIDParams) ([]Reply, error) {
	rows, err := q.db.QueryContext(ctx, getRepliesByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reply
	for rows.Next() {
		var i Reply
		if err := rows.Scan(
			&i.ID,
			&i.TweetID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getReply = `-- name: GetReply :one
SELECT id, tweet_id, user_id, content, created_at, updated_at FROM replies WHERE id = $1
`

func (q *Queries) GetReply(ctx context.Context, id int64) (Reply, error) {
	row := q.db.QueryRowContext(ctx, getReply, id)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.TweetID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateReply = `-- name: UpdateReply :one
UPDATE replies SET
    content = $2,
    updated_at = $3
WHERE id = $1
RETURNING id, tweet_id, user_id, content, created_at, updated_at
`

type UpdateReplyParams struct {
	ID        int64        `json:"id"`
	Content   string       `json:"content"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func (q *Queries) UpdateReply(ctx context.Context, arg UpdateReplyParams) (Reply, error) {
	row := q.db.QueryRowContext(ctx, updateReply, arg.ID, arg.Content, arg.UpdatedAt)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.TweetID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
