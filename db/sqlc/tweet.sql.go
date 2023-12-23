// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: tweet.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTweet = `-- name: CreateTweet :one
INSERT INTO tweets (
    user_id, content, updated_at, retweet_id, n_likes, n_retweets, n_reply
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply
`

type CreateTweetParams struct {
	UserID    int64              `json:"user_id"`
	Content   string             `json:"content"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	RetweetID pgtype.Int8        `json:"retweet_id"`
	NLikes    pgtype.Int4        `json:"n_likes"`
	NRetweets pgtype.Int4        `json:"n_retweets"`
	NReply    pgtype.Int4        `json:"n_reply"`
}

func (q *Queries) CreateTweet(ctx context.Context, arg CreateTweetParams) (Tweet, error) {
	row := q.db.QueryRow(ctx, createTweet,
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
	_, err := q.db.Exec(ctx, deleteTweetByID, id)
	return err
}

const deleteTweetsByUserID = `-- name: DeleteTweetsByUserID :exec
DELETE FROM tweets WHERE user_id = $1
`

func (q *Queries) DeleteTweetsByUserID(ctx context.Context, userID int64) error {
	_, err := q.db.Exec(ctx, deleteTweetsByUserID, userID)
	return err
}

const getTweetByID = `-- name: GetTweetByID :one
SELECT id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply FROM tweets WHERE id = $1
`

func (q *Queries) GetTweetByID(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRow(ctx, getTweetByID, id)
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

const getTweetByUserID = `-- name: GetTweetByUserID :many
SELECT id, user_id, content, created_at, updated_at, is_deleted, retweet_id, n_likes, n_retweets, n_reply FROM tweets WHERE user_id = $1 ORDER BY id LIMIT $2 OFFSET $3
`

type GetTweetByUserIDParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetTweetByUserID(ctx context.Context, arg GetTweetByUserIDParams) ([]Tweet, error) {
	rows, err := q.db.Query(ctx, getTweetByUserID, arg.UserID, arg.Limit, arg.Offset)
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
	rows, err := q.db.Query(ctx, listTweets, arg.Limit, arg.Offset)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
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
	ID        int64              `json:"id"`
	Content   string             `json:"content"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	RetweetID pgtype.Int8        `json:"retweet_id"`
	NLikes    pgtype.Int4        `json:"n_likes"`
	NRetweets pgtype.Int4        `json:"n_retweets"`
	NReply    pgtype.Int4        `json:"n_reply"`
}

func (q *Queries) UpdateTweetByID(ctx context.Context, arg UpdateTweetByIDParams) (Tweet, error) {
	row := q.db.QueryRow(ctx, updateTweetByID,
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